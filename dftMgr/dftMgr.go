package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zededa/zeddp/src/functionTracer/proto/dft"
)

/*type Client struct {
	ClientID	string	`json:"client_id"`
	Jobs		*[]dft.Job	`json:"jobs"`
}

type Job struct {
	ClientID	string	`json:"client_id"`
    JobID		string	`json:"job_id"`
	Tasks		*[]dft.Task	`json:"tasks"`
}

type Task struct {
	TaskID		string	`json:"task_id"`
	Type		string	`json:"type"`
	Repetitions	string	`json:"repetitions"`
	Destination	string	`json:"destination"`
	Timeout		string	`json:"timeout"`
	//StartTime	string	`json:"start_time"`
}

type Result struct {
	ClientID	string	`json:"client_id"`
	JobID		string	`json:"job_id"`
	Results		string	`json:"results"`
}*/

var allJobs map[string]dft.Job       // maps JobID:Job
var allTasks map[string]dft.Task     // maps TaskID:Task
var allResults map[string]dft.Result // maps JobID:Result
var allClients map[string]dft.Client // maps ClientID:Client
var nextJobID, nextTaskID, nextClientID uint64 = 1, 1, 1
var nextJobIDString, nextTaskIDString, nextClientIDString string = "1", "1", "1"

//UI
func createTask(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newTask dft.Task
	_ = json.NewDecoder(req.Body).Decode(&newTask)
	newTask.TaskId = nextTaskIDString
	allTasks[nextTaskIDString] = newTask
	nextTaskID++
	nextTaskIDString = strconv.FormatUint(nextTaskID, 10)
	json.NewEncoder(w).Encode(newTask)
}

func assignTask(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	_, present := allJobs[params["job_id"]]
	if present {
		task := allTasks[params["task_id"]]
		job := allJobs[params["job_id"]]
		job.Tasks = append(job.Tasks, &task)
		allJobs[params["job_id"]] = job
		json.NewEncoder(w).Encode(allJobs[params["job_id"]])
		return
	}
	json.NewEncoder(w).Encode(&dft.Job{})
}

func createJob(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	var newJob dft.Job
	newJob.ClientId = params["client_id"]
	newJob.JobId = nextJobIDString
	newJob.Tasks = []*dft.Task{}
	allJobs[nextJobIDString] = newJob
	nextJobID++
	nextJobIDString = strconv.FormatUint(nextJobID, 10)
	json.NewEncoder(w).Encode(newJob)
}

func createClient(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newClient dft.Client
	newClient.ClientId = nextClientIDString
	newClient.Jobs = []*dft.Job{}
	allClients[nextClientIDString] = newClient
	nextClientID++
	nextClientIDString = strconv.FormatUint(nextClientID, 10)
	json.NewEncoder(w).Encode(newClient)
}

func publishJob(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	_, present := allClients[allJobs[params["job_id"]].ClientId]
	if present {
		client := allClients[allJobs[params["job_id"]].ClientId]
		job := allJobs[params["job_id"]]
		client.Jobs = append(client.Jobs, &job)
		allClients[allJobs[params["job_id"]].ClientId] = client
	} else {
		job := allJobs[params["job_id"]]
		allClients[allJobs[params["job_id"]].ClientId] = dft.Client{allJobs[params["job_id"]].ClientId, []*dft.Job{&job}}
	}
	json.NewEncoder(w).Encode(allJobs[params["job_id"]])
}

func getResult(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	result, present := allResults[params["job_id"]]
	if present {
		json.NewEncoder(w).Encode(result)
		return
	}
	json.NewEncoder(w).Encode(&dft.Result{})
}

//Client
func getJob(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	client, present := allClients[params["client_id"]]
	if present {
		clientJobs := client.Jobs
		if len(clientJobs) > 0 {
			job := clientJobs[0]
			json.NewEncoder(w).Encode(job)
			return
		}
	}
	json.NewEncoder(w).Encode(&dft.Job{})
}

func postResult(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	var newResult dft.Result
	_ = json.NewDecoder(req.Body).Decode(&newResult)
	newResult.JobId = params["job_id"]
	client := allClients[allJobs[params["job_id"]].ClientId]
	client.Jobs = client.Jobs[1:]
	allClients[allJobs[params["job_id"]].ClientId] = client
	delete(allJobs, params["job_id"])
	allResults[params["job_id"]] = newResult
	json.NewEncoder(w).Encode(newResult)
}

//For Testing
func getAllTasks(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allTasks)
}

func getAllClients(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allClients)
}

func getAllJobs(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allJobs)
}

func getAllResults(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allResults)
}

func main() {
	router := mux.NewRouter()

	allJobs = make(map[string]dft.Job)       // maps JobID:Job
	allTasks = make(map[string]dft.Task)     // maps TaskID:Task
	allResults = make(map[string]dft.Result) // maps JobID:Result
	allClients = make(map[string]dft.Client) // maps ClientID:Client

	//UI
	router.HandleFunc("/api/v1/user/create-task", createTask).Methods("POST")
	router.HandleFunc("/api/v1/user/assign-task/{job_id}/{task_id}", assignTask).Methods("POST")
	router.HandleFunc("/api/v1/user/create-job/{client_id}", createJob).Methods("POST")
	router.HandleFunc("/api/v1/user/create-client", createClient).Methods("POST")
	router.HandleFunc("/api/v1/user/publish-job/{job_id}", publishJob).Methods("POST")
	router.HandleFunc("/api/v1/user/get-result/{job_id}", getResult).Methods("GET")

	router.HandleFunc("/api/v1/user/get-all-tasks", getAllTasks).Methods("GET")
	router.HandleFunc("/api/v1/user/get-all-jobs", getAllJobs).Methods("GET")
	router.HandleFunc("/api/v1/user/get-all-results", getAllResults).Methods("GET")
	router.HandleFunc("/api/v1/user/get-all-clients", getAllClients).Methods("GET")

	//Client
	router.HandleFunc("/api/v1/client/get-job/{client_id}", getJob).Methods("GET")
	router.HandleFunc("/api/v1/client/post-result/{job_id}", postResult).Methods("POST")

	log.Fatal(http.ListenAndServe(":12345", router))
}
