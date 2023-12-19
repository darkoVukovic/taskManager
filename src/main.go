package main;

import (
    "fmt"
    "time"
    "encoding/json"
    "os"
    "bufio"
);

type Person struct {
    Id        uint;
    FirstName string;
    Task      string;
    Date      string;
    Status    uint8;
}


func main() {
    fmt.Println("Welcome to task Manger app");

    belgradeLocation, err := time.LoadLocation("Europe/Belgrade");
    if err != nil {
        fmt.Println("timezone not found ", err);
        return;
    }
    currentDate := time.Now().In(belgradeLocation).Format("02-01-2006");
    fmt.Println("current date in belgrade is: ", currentDate);
    // IF selected to create new task 
    for {
        fmt.Println("to create new task press t || to view all tasks press v || to exit press e");
        var input string;
        fmt.Scanln(&input);
        switch input {

        case "t": 
            var id uint = 0;
            var name string;
            var task string;
            var status uint8 = 0;
            // read file first 
            fileContent, err := os.ReadFile("tasks.json");
            if err != nil {
                if os.IsNotExist(err) {
                    fmt.Println("file does not exist. Creating new file ");
                } else {
                    fmt.Println("error loading file ", err);
                    return;
                }
            }
            
            var person []Person;
            err = json.Unmarshal(fileContent, &person);
            if err != nil {
                fmt.Println("Error unmarshaling  file", err);
                return;
            }
            
            if len(person) > 0 {
                lastTask := person[len(person) - 1];
                lastId := lastTask.Id;
                id = lastId + 1;
            }

            
            fmt.Println("enter your name ");
            fmt.Scanln(&name);
            fmt.Println("enter your task");
            scanner := bufio.NewScanner(os.Stdin);
            if scanner.Scan() {
                task = scanner.Text();
            } else {    
                fmt.Println("error reading input ", scanner.Err());
                return;
            }
    
    
            newTask := Person{Id:id, FirstName:name, Task:task, Date:currentDate, Status:status}
            person = append(person, newTask);
            jsonData , err := json.MarshalIndent(person, "", " ");
            if err != nil {
                fmt.Println("error marshaling json", err);
                return;
            }
    
            _ = os.WriteFile("tasks.json", jsonData, os.ModePerm);


        case "v":
            fmt.Println("all tasks ");
            fileContent, err := os.ReadFile("tasks.json");
            if err != nil {
                fmt.Println("Error loading file", err);
                return;
            }
            var person []Person;
            err = json.Unmarshal(fileContent, &person);
            if err != nil {
                fmt.Println("error unmarshaling JSON", err);
                return;
            }
    
            fmt.Println("tasks from file :");
            for _, person := range person {
                fmt.Printf("ID: %d, Name: %s, Task: %s, Date:%s ,Status: %d \n", person.Id, person.FirstName, person.Task,person.Date,  person.Status);
            }
        case "e" :
            fmt.Println("exiting app");
            return;

        default: 
            fmt.Println("invalid command");
        }

  
    }


}   