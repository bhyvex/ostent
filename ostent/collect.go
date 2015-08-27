package ostent

import (
	"os"
	"regexp"
	"runtime"
	"strings"
	"sync"

	"github.com/ostrost/ostent/format"
	"github.com/ostrost/ostent/system"
	"github.com/ostrost/ostent/system/operating"
	sigar "github.com/rzab/gosigar"
)

type IfData struct {
	IP         string
	Name       string
	InBytes    uint
	OutBytes   uint
	InPackets  uint
	OutPackets uint
	InErrors   uint
	OutErrors  uint
}

func (id IfData) GetInBytes() uint    { return id.InBytes }
func (id IfData) GetOutBytes() uint   { return id.OutBytes }
func (id IfData) GetInErrors() uint   { return id.InErrors }
func (id IfData) GetOutErrors() uint  { return id.OutErrors }
func (id IfData) GetInPackets() uint  { return id.InPackets }
func (id IfData) GetOutPackets() uint { return id.OutPackets }

// Registry has updates with sigar values.
type Registry interface {
	UpdateIFdata(IfData)
	UpdateCPU([]sigar.Cpu)
	UpdateLoadAverage(sigar.LoadAverage)
	UpdateSwap(sigar.Swap)
	UpdateRAM(sigar.Mem, uint64, uint64)
	UpdateDF(sigar.FileSystem, sigar.FileSystemUsage)
}

// Collector is collection interface.
type Collector interface {
	GetHostname() (string, error)
	Hostname(S2SRegistry, *sync.WaitGroup)
	Uptime(S2SRegistry, *sync.WaitGroup)
	LA(Registry, *sync.WaitGroup)
	RAM(Registry, *sync.WaitGroup)
	Swap(Registry, *sync.WaitGroup)
	Interfaces(Registry, S2SRegistry, *sync.WaitGroup)
	Procs(chan<- ProcSlice)
	Disks(Registry, *sync.WaitGroup)
	CPU(Registry, *sync.WaitGroup)
}

var (
	// RXlo is a regexp to match loopback network interface
	RXlo = regexp.MustCompile("^lo\\d*$")

	// RXfw is a regexp to match non-hardware network interface
	RXfw = regexp.MustCompile("^fw\\d+$")
	// RXgif is a regexp to match non-hardware network interface
	RXgif = regexp.MustCompile("^gif\\d+$")
	// RXstf is a regexp to match non-hardware network interface
	RXstf = regexp.MustCompile("^stf\\d+$")
	// RXwdl is a regexp to match non-hardware network interface
	RXwdl = regexp.MustCompile("^awdl\\d+$")
	// RXbridge is a regexp to match non-hardware network interface
	RXbridge = regexp.MustCompile("^bridge\\d+$")
	// RXvboxnet is a regexp to match non-hardware network interface
	RXvboxnet = regexp.MustCompile("^vboxnet\\d+$")
	// RXairdrop is a regexp to match non-hardware network interface
	RXairdrop = regexp.MustCompile("^p2p\\d+$")
)

// HardwareInterface returns false for known virtual/software network interface name.
func HardwareInterface(name string) bool {
	if RXbridge.MatchString(name) ||
		RXvboxnet.MatchString(name) {
		return false
	}
	if runtime.GOOS == "darwin" {
		if RXfw.MatchString(name) ||
			RXgif.MatchString(name) ||
			RXstf.MatchString(name) ||
			RXwdl.MatchString(name) ||
			RXairdrop.MatchString(name) {
			return false
		}
	}
	return true
}

// Machine implements Collector by collecting the maching metrics.
type Machine struct{}

func (m Machine) GetHostname() (string, error) {
	// m is unused
	return GetHostname()
}

func GetHostname() (string, error) {
	hostname, err := os.Hostname()
	if err == nil {
		hostname = strings.Split(hostname, ".")[0]
	}
	return hostname, err
}

func (m Machine) Hostname(sreg S2SRegistry, wg *sync.WaitGroup) {
	if hostname, err := m.GetHostname(); err == nil {
		sreg.SetString("hostname", hostname)
	}
	wg.Done()
}

func (m Machine) Uptime(sreg S2SRegistry, wg *sync.WaitGroup) {
	// m is unused
	uptime := sigar.Uptime{}
	uptime.Get()
	sreg.SetString("uptime", format.FormatUptime(uptime.Length))
	wg.Done()
}

func (m Machine) LA(reg Registry, wg *sync.WaitGroup) {
	// m is unused
	la := sigar.LoadAverage{}
	la.Get()
	reg.UpdateLoadAverage(la)
	wg.Done()
}

func _getmem(kind string, in sigar.Swap) operating.Memory {
	total, approxtotal, _ := format.HumanBandback(in.Total)
	used, approxused, _ := format.HumanBandback(in.Used)

	return operating.Memory{
		Kind:       kind,
		Total:      total,
		Used:       used,
		Free:       format.HumanB(in.Free),
		UsePercent: format.FormatPercent(approxused, approxtotal),
	}
}

func (m Machine) RAM(reg Registry, wg *sync.WaitGroup) {
	// m is unused
	got := sigar.Mem{}
	extra1, extra2, _ := sigar.GetExtra(&got)
	reg.UpdateRAM(got, extra1, extra2)
	wg.Done()

	// inactive := got.ActualFree - got.Free // == got.Used - got.ActualUsed // "kern"
	// _ = inactive

	// Used = .Total - .Free
	// | Free |           Used +%         | Total
	// | Free | Inactive | Active | Wired | Total

	// TODO active := vm_data.active_count << 12 (pagesize)
	// TODO wired  := vm_data.wire_count   << 12 (pagesoze)
}

func (m Machine) Swap(reg Registry, wg *sync.WaitGroup) {
	// m is unused
	got := sigar.Swap{}
	got.Get()
	reg.UpdateSwap(got)
	wg.Done()
}

func (m Machine) Disks(reg Registry, wg *sync.WaitGroup) {
	// m is unused
	fls := sigar.FileSystemList{}
	fls.Get()

	// devnames := map[string]bool{}
	dirnames := map[string]bool{}

	for _, fs := range fls.List {

		usage := sigar.FileSystemUsage{}
		usage.Get(fs.DirName)

		if !strings.HasPrefix(fs.DevName, "/") {
			continue
		}
		// if _, ok := devnames[fs.DevName]; ok
		if _, ok := dirnames[fs.DirName]; ok {
			continue
		}
		// devnames[fs.DevName] = true
		dirnames[fs.DirName] = true

		reg.UpdateDF(fs, usage)
	}
	wg.Done()
}

func (m Machine) Procs(CH chan<- ProcSlice) {
	// m is unused
	var procs ProcSlice
	pls := sigar.ProcList{}
	pls.Get()

	for _, pid := range pls.List {

		state := sigar.ProcState{}
		// args := sigar.ProcArgs{}
		time := sigar.ProcTime{}
		mem := sigar.ProcMem{}

		if err := state.Get(pid); err != nil {
			continue
		}
		// if err :=  args.Get(pid); err != nil { continue }
		if err := time.Get(pid); err != nil {
			continue
		}
		if err := mem.Get(pid); err != nil {
			continue
		}

		procs = append(procs,
			operating.ProcInfo{
				PID:      uint(pid),
				Priority: state.Priority,
				Nice:     state.Nice,
				Time:     time.Total,
				Name:     system.ProcName(pid, state.Name),
				// Name:  strings.Join(append([]string{system.ProcName(pid, state.Name)}, args.List[1:]...), " "),
				UID:      state.Uid,
				Size:     mem.Size,
				Resident: mem.Resident,
			})
	}
	CH <- procs
}

func (m Machine) CPU(reg Registry, wg *sync.WaitGroup) {
	// m is unused
	cl := sigar.CpuList{}
	cl.Get()
	reg.UpdateCPU(cl.List)
	wg.Done()
}
