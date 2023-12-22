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
    // TODO updates and deletions(loading file once on top or on every case ??) apply DRY later
    for {
        fmt.Println("to create new task press t || to view all tasks press v || to delete task press d || press u to update task ||  to exit press e");
        var input string;
        const file string =  "tasks.json";
        fmt.Scanln(&input);
        switch input {
    // IF selected to create new task 
        case "t": 
            var id uint = 0;
            var name string;
            var task string;
            var status uint8 = 0;
            // read file first 
            fileContent, err := os.ReadFile(file);
            if err != nil {
                if os.IsNotExist(err) {
                    fmt.Println("file does not exist. Creating new file ");
                    fileContent, createErr := os.Create(file);
                    if createErr != nil {
                        fmt.Println("error creating file", createErr);
                        return;
                    }
                    defer fileContent.Close();
                    
                } else {
                    fmt.Println("error loading file ", err);
                    return;
                }
            }
            
            var person []Person;
            
            if(len(fileContent) > 0) {
                err = json.Unmarshal(fileContent, &person);
                if err != nil {
                    fmt.Println("Error unmarshaling  file", err);
                    return;
                }
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
    
            _ = os.WriteFile(file, jsonData, os.ModePerm);

        // display all tasks 
        // TODO display tasks for specicc user(finished)
        case "v":
            ViewTasks(file);
        case "e" :
            fmt.Println("exiting app.");
            return;
        case "d": 
            var deleteId uint;
            ViewTasks(file);
            fmt.Println("choose id of task to delete:");
            fmt.Scanln(&deleteId);
            DeleteTask(file, deleteId);
        case "u":
            var updateId uint;
            var status uint8;
            ViewTasks(file);
            fmt.Println("choose id of task to update");
            fmt.Scanln(&updateId);
            fmt.Println("change status to solved(1) or not solved(0)");
            fmt.Scanln(&status);
            if !(status == 0 || status == 1) {
                fmt.Println("invalid status. Type 0 for not solved or 1 for solved tasks");
            } else {
                UpdateTask(file, updateId, status);
            }
        default: 
            fmt.Println("invalid command");
        }

  
    }


}   


    // view tasks for specific username 
    // TODO make it aslo delete or make another func for it 
    func ViewTasks(filename string) {
        var name string;
        var exist bool = false;
        fmt.Println("your username: ");
        fmt.Scanln(&name);
        fileContent, err := os.ReadFile(filename);
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
        for _, person := range person {
            if person.FirstName == name {
                fmt.Printf("ID: %d, Name: %s, Task: %s, Date:%s ,Status: %d \n", person.Id, person.FirstName, person.Task,person.Date,  person.Status);
                exist = true;
            }
        }
        if !exist {
            fmt.Println("user does not have any tasks.");
            return;
        }
    }

    func DeleteTask(filename string, deleteId uint) {
        fileContent, err := os.ReadFile(filename);
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
        for i, row := range person {
            if row.Id == deleteId {
                person = append(person[:i], person[i+1:]...)
                break;
            }
        }
        jsonData , err := json.MarshalIndent(person, "", " ");
        if err != nil {
            fmt.Println("error marshaling json", err);
            return;
        }

        _ = os.WriteFile(filename, jsonData, os.ModePerm);
    } 
    
    func UpdateTask(filename string, updateId uint, status uint8) {
        fileContent, err := os.ReadFile(filename);
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
        for i, row := range person {
            if row.Id == updateId {
                person[i].Status = status;
                break;
            }
        }
        jsonData , err := json.MarshalIndent(person, "", " ");
        if err != nil {
            fmt.Println("error marshaling json", err);
            return;
        }

        _ = os.WriteFile(filename, jsonData, os.ModePerm);
        fmt.Println("update completed on id", updateId);
    } 


