package main

import(
  "os"
  "log"
  "bufio"
  "time"
  "net"
  "fmt"

  "github.com/codegangsta/cli"
)

  var message = "Usage: gomsg send <file> <ip:port>"
  var cfmessage = "Usage: gomsg cf <syslog up:port> <graphite ip:port>"

func usage(){
      log.Fatalf(message)
}

func main(){
  app := cli.NewApp()
  app.Name = "gomsg"
  app.Usage = "gomsg - send logs from a file"
  app.Version = "1.0"
  app.Action = func(c *cli.Context) {
    cli.ShowAppHelp(c)
  }
  app.Commands = []cli.Command{
    {
      Name:       "send",
      ShortName:  "s",
      Usage:      message,
      Action: func(c *cli.Context){
        args := c.Args()
        if len(args) < 2 {
          cli.ShowAppHelp(c)
          os.Exit(1)
        }

        var logs []string;
        logs = parseFile(args[0])
        var server = args[1]

        fmt.Println("\nConnected and sending Logs")
        for j := 0; j < 50; j++ {
          time.Sleep(5 * time.Second)
          for i := 0; i < len(logs); i++ {
            time.Sleep(500 * time.Millisecond)
            sendLogs(logs[i], server)
          }
        }

      },
    },
    {
      Name:       "cf",
      Usage:      cfmessage,
      Action: func(c *cli.Context){
        args := c.Args()
        if len(args) < 2 {
          cli.ShowAppHelp(c)
          os.Exit(1)
        }

        var syslogs []string;
        var graphite []string;
        syslogs = parseFile("syslog.txt")
        graphite = parseFile("graphite.txt")

        fmt.Println("\nConnected and sending Logs")
        for j := 0; j < 50; j++ {
          time.Sleep(5 * time.Second)
          for i := 0; i < len(syslogs); i++ {
            time.Sleep(350 * time.Millisecond)
            sendLogs(graphite[i], args[1])
            sendLogs(syslogs[i], args[0])
          }
        }

      },
    },
  }
  app.Run(os.Args)
}

func parseFile(f string) ([]string){
  var logs []string;
  file, err := os.Open(f)
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()
  scanner := bufio.NewScanner(file)
  for scanner.Scan(){
    logs = append(logs, scanner.Text())
  }
  if err := scanner.Err(); err != nil{
    log.Fatal(err)
  }

  return logs
}

func sendLogs(logs string, s string){
  conn, err := net.Dial("tcp", s)
  if err != nil {
    log.Fatal(err)
    os.Exit(1)
  }
  if len(logs) > 1 {
    writer := bufio.NewWriter(conn)
    writer.WriteString(logs)
    writer.Flush()
  }

  conn.Close()
}
