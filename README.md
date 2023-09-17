# Task Manager CLI

The Task Manager CLI is a simple command-line application that allows you to manage your tasks. You can add, list, mark as completed, and remove tasks using this application.

## Usage

1. **Adding a Task**
   - To add a new task, select option "1" from the menu.
   - Enter the task title when prompted.
   - Enter a task description when prompted.
   - Enter the due date in the format "YYYY-MM-DD" when prompted.

2. **Listing Tasks**
   - To list all tasks, select option "2" from the menu.
   - Tasks will be displayed with their ID, Title, Description, Due Date, and Status (Completed or Pending).

3. **Marking a Task as Completed**
   - To mark a task as completed, select option "3" from the menu.
   - Enter the ID of the task you want to mark as completed when prompted.
   - The task's status will be updated to "Completed."

4. **Removing a Task**
   - To remove a task, select option "4" from the menu.
   - Enter the ID of the task you want to remove when prompted.
   - The task will be permanently removed from the list.

5. **Exiting the Application**
   - To exit the application, select option "5" from the menu.
   - Your task data will be saved to a file, so you can continue where you left off next time.

## Installation

1. Clone the repository to your local machine:

    ```bash
   git clone https://github.com/marcovoliveira/go-task-manager-cli.git
    ```

2. Navigate to the project directory:

     ```bash
   cd your-task-manager-cli
    ```
3. Run the application:

    ```bash
   go run main.go
    ```
## File Persistence

Your task data is automatically saved to a CSV file named "tasks.csv" in the project directory. This allows you to continue managing your tasks across different sessions.

## Example Task Format

Tasks are stored in the CSV file in the following format:
```csv
ID,Title,Description,Due Date,Completed
1,"Finish presentation","Prepare slides for the upcoming meeting",2023-09-20,false
2,"Go for a run","Run 5 miles in the park",2023-09-25,false
```

**Task Format:**

- **ID**: Unique identifier for the task.
- **Title**: A brief, descriptive title for the task.
- **Description**: Additional details or instructions for the task.
- **Due Date**: The date by which the task should be completed (in the format YYYY-MM-DD).
- **Status**: Indicates whether the task is completed or pending (true for completed, false for pending).

**Example Task:**

- **ID**: 1
- **Title**: Finish presentation
- **Description**: Prepare slides for the upcoming meeting
- **Due Date**: 2023-09-20
- **Status**: false

## License

This project is licensed under the [GNU General Public License v3.0](LICENSE.md) - see the [LICENSE.md](LICENSE.md) file for details.
