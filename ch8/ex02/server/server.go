package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const CONNECT_PORT = 8000

func main() {
	listener, err := net.Listen("tcp", "localhost:"+strconv.Itoa(CONNECT_PORT))
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go dataConn(conn)
	}
}

func dataConn(c net.Conn) {
	defer c.Close()
	fmt.Fprintln(c, "220 Welcome to my FTP-server")
	input := bufio.NewScanner(c)

	var conn net.Conn

	for input.Scan() {
		texts := strings.Split(input.Text(), " ")
		fmt.Println(texts)
		if len(texts) > 0 {
			switch cmd := strings.ToUpper(texts[0]); cmd {
			case "LIST":
				err := handleList(c, texts)
				if err != nil {
					fmt.Fprintln(c, "400")
					break
				}
				fmt.Fprintf(c, "200")
			case "RETR": // get
				err := handleRetr(conn, texts)
				if err != nil {
					fmt.Fprintln(c, "400")
					break
				}
				fmt.Fprintln(c, "200 RETR")
			case "STOR": // put
				fmt.Fprintln(c, "150 STOR")
				err := handleStor(conn, texts)
				if err != nil {
					fmt.Fprintln(c, "400")
					break
				}
				fmt.Fprintln(c, "200 STOR")
			case "EXIT":
				fmt.Fprintln(c, "Close connection")
				c.Close()
				if conn != nil {
					conn.Close()
				}
				return
			case "PORT":
				if len(texts) > 1 {
					conn = handlePort(texts[1])
					if conn == nil {
						fmt.Fprintln(c, "400")
						break
					}
					fmt.Fprintln(c, "200")
				}
			case "USER":
				fmt.Fprintln(c, "230")
			case "SYST":
				fmt.Fprintln(c, "215")
			case "QUIT":
				fmt.Fprintln(c, "Close connection")
				c.Close()
				if conn != nil {
					conn.Close()
				}
				return
			default:
				fmt.Fprintln(c, "501 Command not found")
			}
		}
	}
}

func handlePort(arg string) net.Conn {
	fmt.Println(arg)
	texts := strings.Split(arg, ",")
	if len(texts) != 6 {
		return nil
	}
	clientAddr := texts[0] + "." + texts[1] + "." + texts[2] + "." + texts[3]
	p1, _ := strconv.Atoi(texts[4])
	p2, _ := strconv.Atoi(texts[5])

	clientPort := p1*256 + p2
	conn, err := net.Dial("tcp", clientAddr+":"+strconv.Itoa(clientPort))
	if err != nil {
		return nil
	}

	return conn
}

func handleStor(conn io.Reader, texts []string) error {
	var w io.Writer
	if len(texts) > 1 {
		f, err := os.Create("./" + texts[1])
		if err != nil {
			return err
		}

		defer f.Close()
		w = bufio.NewWriter(f)
	} else {
		w = new(bytes.Buffer)
	}

	if _, err := io.Copy(w, conn); err != nil {
		log.Fatal(err)
	}
	return nil
}

func handleList(conn net.Conn, texts []string) error {
	path := "./"
	if len(texts) > 1 {
		path = texts[1]
	}
	cmd := exec.Command("ls", path)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return err
	}
	results := strings.Replace(out.String(), "\n", "\t", -1)
	fmt.Fprintln(conn, results)
	return nil
}

func handleRetr(conn net.Conn, texts []string) error {
	if len(texts) < 2 {
		return fmt.Errorf("no filepath")
	}

	path := texts[1]

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	_, err = conn.Write(content)
	if err != nil {
		return err
	}

	return nil
}
