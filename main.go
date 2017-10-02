package main

import (
  "fmt"
  "log"
  "os"
  "net/http"
  "github.com/stianeikeland/go-rpio"
  "github.com/gorilla/mux"
  "encoding/json"
)

var dc1_a, dc1_b, dc2_a, dc2_b, led rpio.Pin
const interval = 30
var lit bool

type Motion struct {
	Direction string `json:"direction"`
}

func main(){
        rpio.Open()
        defer rpio.Close()
        Initiate()

        router := mux.NewRouter().StrictSlash(true)
	      router.HandleFunc("/", index)
		    router.HandleFunc("/api/motion", motion_controller)
        router.HandleFunc("/api/lights", lights_controller)
		    log.Fatal(http.ListenAndServe(":80", router))
}

//Handler for JSON calls
func motion_controller(w http.ResponseWriter, r *http.Request) {
  	switch r.Method {
  	default:
  		w.WriteHeader(http.StatusMethodNotAllowed)
  		fmt.Fprintln(w, "Invalid request method")

  	case "POST":
  		var mtn Motion
  		dec := json.NewDecoder(r.Body)
  		err := dec.Decode(&mtn)
  		if err != nil {
  			w.WriteHeader(http.StatusBadRequest)
  			fmt.Fprintf(w, "Invalid JSON request")
  		} else {
  			switch{
  			case mtn.Direction == "forward"
          Forward(interval)
    			fmt.Fprintf(w, "ok")

        case mtn.Direction == "left":
          Left(interval)
          fmt.Fprintf(w, "ok")

        case mtn.Direction == "right":
          Right(interval)
          fmt.Fprintf(w, "ok")

        case mtn.Direction == "backward":
          Backwards(interval)
          fmt.Fprintf(w, "ok")

  			default:
          w.WriteHeader(http.StatusBadRequest)
    			fmt.Fprintf(w, "Invalid direction")
  		}
    }
  }
}

func lights_controller(w http.ResponseWriter, r *http.Request) {
  	switch r.Method {
  	default:
  		w.WriteHeader(http.StatusMethodNotAllowed)
  		fmt.Fprintln(w, "Invalid request method")

  	case "GET":
        if lit {
          Lightsoff()
          fmt.Fprintf(w, "LEDs disabled")
        }else {
          Lightson()
          fmt.Fprintf(w, "LEDs enabled")
        }
  			fmt.Fprintf(w, "ok")
  		}
    }
  }
}
