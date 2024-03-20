# guess-it-1

## Usage

### Cloning the Repository

```bash
git clone git@git.01.alem.school:azhaxylyk/guess-it-1.git
```
### Setting up the Environment

Download [zip file](https://assets.01-edu.org/guess-it/guess-it-dockerized.zip) and extract it.

Once you have the repository locally, navigate into the guess-it-dockerized/ directory. You will need to place the student/ folder (provided by the student) inside the root directory. This folder contains the student's guessing program along with a file named script.sh, which should be an executable shell script that runs the student's program. Your file system should resemble the following structure:

```console
─ guess-it-dockerized/
├── ai/
│   ├── big-range
│   └── ...
├── index.html
├── index.js
└── ...
└── student/
    ├── solution.py
    └── script.sh

```

Before running the program, ensure that the script.sh file is executable. You can make it executable using the following command in your terminal:

```console
chmod +x student/script.sh
```

### Testing the Student Program

1. Ensure you have Docker installed on your system.
2. Run the following commands to set up the dependencies and start the webpage on port 3000:

```console
docker-compose up
```

3. Open your preferred web browser and go to http://localhost:3000/.
4. Click on any of the "Test Data" buttons. You will notice a message in the Dev Tools/Console indicating that you need another guesser besides the student.
5. To add a guesser, append the guesser's name to the URL. The guesser's name should match one of the files present in the ai/ folder. For example

```console
http://localhost:3000/?guesser=big-range
```

6. Choose which random data set to test. You can either wait for the program to test all values or click "Quick" to skip the waiting and see the results immediately.
7. After each test, it's recommended to clear the displays by clicking the "Clean" button.

### Additional Notes
- The website utilizes large data sets, so it's advisable to clear the displays after each test for better performance.
- [Audit](https://github.com/01-edu/public/tree/master/subjects/guess-it-1/audit)