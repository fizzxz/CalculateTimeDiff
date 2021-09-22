package taskFunctions

import "go.mongodb.org/mongo-driver/bson/primitive"

type ToDoTask struct {
	TaskID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`

	TaskCategory    string `json:"category,omitempty"`
	TaskStatus      bool   `json:"status,omitempty"`
	TaskDescription string `json:"description,omitempty"`
	TaskPriority    int    `json:"priority,omitempty"`
	TaskStartDate   string `json:"startDate,omitempty"`
	TaskDueDate     string `json:"dueDate,omitempty"`
}

func NewToDoTask(taskID primitive.ObjectID, taskCategory, taskDescription, taskStartDate,
	taskDueDate string, taskPriority int, taskStatus bool) *ToDoTask {
	return &ToDoTask{
		TaskID:          taskID,
		TaskCategory:    taskCategory,
		TaskStatus:      taskStatus,
		TaskDescription: taskDescription,
		TaskPriority:    taskPriority,
		TaskStartDate:   taskStartDate,
		TaskDueDate:     taskDueDate,
	}
}

type taskFunctions interface {
	Close() error
	InsertTask(user *ToDoTask)
}
