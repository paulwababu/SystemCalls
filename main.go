/*
The Go program below does the following basic operation:
                                  1. File Manipulation.
                                  2. Device management.(check open ports on server(both local and remote))
                                  3. Memory management.
                                  4. Time management(show current time and future time even generations from now)
                                  5. Process management(Start(open) mozilla firefox and,spawn and kill process using exec command)
                                    *6. EXTRA OPTION FOR PORT CHECKING TOOL*
The Program also has options of scanning ports on remote servers, kindly do not use it for illegal purposes.
The ip scanner comes in two, one of them is a basic port scanner
therefore not as fast and efficient and the other I have used a library called Furious ID/Port Scanner.
The difference between the two is that furious is fast, lightweight and can be deployed on a remote server.
You would be required to install winPcap, on windows, on linux install libpcap with your package manager,
on OSX, brew install libpcap
*/
package main

import (
	"fmt"
	"github.com/anvie/port-scanner"
	"github.com/mackerelio/go-osstat/memory"
	"github.com/uniplaces/carbon"
	"log"
	"os"
	"os/exec"
	"runtime"
	"context"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
	"golang.org/x/sync/semaphore"
)

func main()  {
	//greats user and asks them the kind of basic operation they would like to perform
	var name string
	fmt.Println("What is your name? ")
	fmt.Scanln(&name)
	fmt.Println("Hello", name, ". Enjoy the code! \n")
	fmt.Println("Basic Operation done by program: \n")
	fmt.Println("1. File manipulation\n" +
		"2. Device management by checking open ports on both local host server and remote server\n" +
		"3. Memory management by displaying memory status \n" +
		"4. Time management by displaying current time and Future time and past time\n" +
		"5. Process management by opening a webpage. And Opening an application on PC ")


	//tell user to pick an operation
	var input1 int
	fmt.Println("Choose one: \n")
	fmt.Scanln(&input1)

	//check user input scanned above this comment and decide on operation to do
	if input1 == 1 {
		fmt.Println("Basic file manipulation: \n" +
			"1. Create New File \n" +
			"2. Rename file name \n" +
			"3. Get file information \n" +
			"4. Delete a file")
		var fileAction int
		fmt.Println("Choose one action: ")
		fmt.Scanln(&fileAction)
		switch fileAction {
		case 1:
			one()
			fmt.Println("Created file!Check in current directory")
		case 2:
			two()
			fmt.Println("File renamed, if not then create a file first")
		case 3:
			three()
		case 4:
			four()
			fmt.Println("Specified file deleted successfully! You can confirm!")
		}
	}
//Device management functions
	if input1 == 2 {
		ps := &PortScanner{
			ip:   "127.0.0.1",
			lock: semaphore.NewWeighted(Ulimit()),
		}
		ps.Start(1, 65535, 100*time.Millisecond)
	}
//memory management functions
	if input1 == 3 {
		eight()
	}
//time management functions
	if input1 == 4 {
		nine()
	}
//process management
    if input1 == 5 {
    	fmt.Println("Here a user can either open VirtualBox application, a website ie wikipaedia, and also spawn and kill a process using system commands \n")
    	var opn int
    	fmt.Println("To open VirtualBox, press 1 \n" +
    		"To open wikipedia from Mozilla firefox, press 2 \n" +
    		"To spawn a new process and kill it, press 3 \n")
    	fmt.Scanln(&opn)
		switch opn {
		case 1:
			ten()
		case 2:
			eleven()
		case 3:
			twelve()
		default:
			fmt.Println("error!")
		}
	}
}

//function to create an empty file
func one()  {
	//declare and initiate variable to create empty file and return
	emptyFile, err := os.Create("Sample_1.txt")
	if err!= nil {
		log.Fatal(err)
	}
	log.Println(emptyFile)
	emptyFile.Close()
}

//function to rename a file
func two()  {
	oldName := "Sample_1.txt"
	newName := "Sample_2.txt"
	err := os.Rename(oldName, newName)
	if err != nil {
		log.Fatal(err)
	}
}

//function to get file information 
func three()  {
	fileStat, err := os.Stat("Sample_1.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File Name:", fileStat.Name())        // Base name of the file
	fmt.Println("Size:", fileStat.Size())             // Length in bytes for regular files
	fmt.Println("Permissions:", fileStat.Mode())      // File mode bits
	fmt.Println("Last Modified:", fileStat.ModTime()) // Last modification time
	fmt.Println("Is Directory: ", fileStat.IsDir())
}

//function to delete a specific file
func four()  {
	err := os.Remove("C:/Users/Paulkiragu621/go/src/SystemCalls/Sample_1.txt")
	if err != nil {
		log.Fatal(err)
	}
}

//function to display os system running
func five()  {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

//function to check open ports on local host server using "github.com/anvie/port-scanner"
func six()  {
	// scan localhost with a 2 second timeout per port in 50 concurrent threads
	ps := portscanner.NewPortScanner("localhost", 2*time.Second, 50)

	// get opened port
	fmt.Printf("scanning port %d-%d...\n", 20, 30000)

	openedPorts := ps.GetOpenedPort(20, 30000)

	for i := 0; i < len(openedPorts); i++ {
		port := openedPorts[i]
		fmt.Print(" ", port, " [open]")
		fmt.Println("  -->  ", ps.DescribePort(port))
	}
}


//function to display memory using "github.com/mackerelio/go-osstat/memory" package
func eight()  {
	memory, err := memory.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
	fmt.Printf("memory total: %d bytes\n", memory.Total)
	fmt.Printf("memory used: %d bytes\n", memory.Used)
	fmt.Printf("memory free: %d bytes\n", memory.Free)
}

//function display current and future time
func nine()  {
	//display current time
	fmt.Printf("Right now is %s\n", carbon.Now().DateTimeString())
	today, _ := carbon.NowInLocation("Africa/Nairobi")
	fmt.Printf("Right now in Nairobi is %s\n", today)
    
	//display any time on calender regardless of period
	fmt.Printf("\n#################################################################\n")
	fmt.Printf("Tomorrow is %s\n", carbon.Now().AddDay())
	fmt.Printf("Yesterday is %s\n", carbon.Now().SubDay())
	fmt.Printf("After 5 days %s\n", carbon.Now().AddDays(5))
	fmt.Printf("Before 5 days %s\n", carbon.Now().SubDays(5))
	fmt.Printf("#################################################################\n")

	fmt.Printf("\n################################################################\n")
	fmt.Printf("Next week is %s\n", carbon.Now().AddWeek())
	fmt.Printf("5 Weeks from now %s\n", carbon.Now().AddWeeks(5))
	fmt.Printf("5 weeks before %s\n", carbon.Now().SubWeeks(5))
	fmt.Printf("\n##################################################################\n")


	fmt.Printf("\n#####################################################################\n")
	fmt.Printf("Next century is %s\n", carbon.Now().AddCentury())
	fmt.Printf("50000 centuries from now is %s\n", carbon.Now().AddCenturies(50000))
	fmt.Printf("After 100 years from now is %s\n", carbon.Now().AddYears(100))
	fmt.Printf("Before 22 years from now it was %s\n", carbon.Now().SubYears(22))
	fmt.Printf("3000 years before now was %s\n", carbon.Now().SubYears(3000))
	fmt.Printf("\n######################################################################\n")

	fmt.Printf("\n#######################################\n")
	fmt.Printf("Weekday? %t\n", carbon.Now().IsWeekday())
	fmt.Printf("Weekend? %t\n", carbon.Now().IsWeekend())
	fmt.Printf("LeapYear? %t\n", carbon.Now().IsLeapYear())
	fmt.Printf("Past? %t\n", carbon.Now().IsPast())
	fmt.Printf("Future? %t\n", carbon.Now().IsFuture())

	fmt.Printf("\n#######################################\n")
	fmt.Printf("Start of day:   %s\n", today.StartOfDay())
	fmt.Printf("End of day: %s\n", today.EndOfDay())
	fmt.Printf("Start of month: %s\n", today.StartOfMonth())
	fmt.Printf("End of month:   %s\n", today.EndOfMonth())
	fmt.Printf("Start of year:  %s\n", today.StartOfYear())
	fmt.Printf("End of year:    %s\n", today.EndOfYear())
	fmt.Printf("Start of week:  %s\n", today.StartOfWeek())
	fmt.Printf("End of week:    %s\n", today.EndOfWeek())
	fmt.Printf("\n##############################################\n")
}

//function to open VMBox which is an example of process management
func ten()  {
	// Create command.
	// ... The application is not executed yet.
	cmd := exec.Command("C:/Program Files/Oracle/VirtualBox/VirtualBox")
	fmt.Println("Starting command")
	// Run virtual box.
	// ... This starts a virtual box instance.
	//     It does not wait for it to finish.
	cmd.Start()
	fmt.Println("DONE")
}

//function to open mozilla firefox and go to a specified url
func eleven()  {
	browser := "C:/ProgramData/Microsoft/Windows/Start Menu/Programs/Google Chrome Beta"
	argument := "https://en.wikipedia.org/";
	cmd := exec.Command(browser, argument)

	// Run firefox and load URL specified in argument.
	cmd.Start()
	fmt.Println("DONE")
}

//function to spwan a new process
//call cmd.Process.Kill() when your condition triggers:
func twelve()  {
	cmd := exec.Command("Coding is life!!!", "5")
	start := time.Now()
	time.AfterFunc(3*time.Second, func() { cmd.Process.Kill() })
	err := cmd.Run()
	fmt.Printf("pid=%d duration=%s err=%s\n", cmd.Process.Pid, time.Since(start), err)
}

//port scanner struct with 2 fields
type PortScanner struct {
	ip string
	lock *semaphore.Weighted
}
//The ip will be the IP address of the host on the network we’re interested in scanning.
// The lock will act as a threshold that will limit the number of go routines(lightweight thread of execution)
// that will be running at any given time.
//To set our lock based on the limits of the operating system, we may opt to use the value of the ulimit command.
// This can be captured to be parsed using the os/exec package.

//ulimit is a builtin shell command used to manage various resource restrictions.
// The -n flag will reveal the maximum number of open files allowed;
// we can capture the output of this command and convert it to the right type.


//To abstract that whole process away of getting the value from the command-line, I decided to put it in a function called Ulimit.
func Ulimit() int64 {
	out, err := exec.Command("ulimit", "-n").Output()
	if err != nil {
		panic(err)
	}
	s := strings.TrimSpace(string(out))
	i, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		panic(err)
	}
	return i
}

//port scanner function using net package
func ScanPort(ip string, port int, timeout time.Duration) {
	target := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", target, timeout)

	if err != nil {
		if strings.Contains(err.Error(), "too many open files") {
			time.Sleep(timeout)
			ScanPort(ip, port, timeout)
		} else {
			fmt.Println(port, "closed")
		}
		return
	}

	conn.Close()
	fmt.Println(port, "open")
}

//PortScanner struct that will scan a given range of ports (from the first f to the last one l)
// along with a Duration to use as a timeout
func (ps *PortScanner) Start(f, l int, timeout time.Duration) {
	wg := sync.WaitGroup{}
	defer wg.Wait()

	for port := f; port <= l; port++ {
		wg.Add(1)
		ps.lock.Acquire(context.TODO(), 1)
		go func(port int) {
			defer ps.lock.Release(1)
			defer wg.Done()
			ScanPort(ps.ip, port, timeout)
		}(port)
	}
}


