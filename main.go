package mylogger

import "log"

func LogInfo(message string){
  log.Printf("info %v",message)
}
