#Script to build all the solutions for different days
import os, subprocess

def main():
    #Go two directories up
    #Get a list of all directories, sequentially go into them and run go build
    current_dir = os.getcwd()
    main_dir = os.path.join(current_dir,os.pardir)
    os.chdir(main_dir)
    directories = [d for d in os.listdir(main_dir) if os.path.isdir(os.path.join(main_dir, d))]
    print(directories)
    print('*' * 20)
    for directory in directories:
        current_dir = os.getcwd()
        path = os.path.join(main_dir, directory)
        all_files = os.listdir(os.getcwd())
        go_files = [f for f in  all_files if os.path.isfile(os.path.join(current_dir, f)) and f.endswith(".go")]
        #Build these golang files
        build_code()

def build_code():
    #Build all the golang files in current directory
    # Run 'go build' command to build Golang code
    process = subprocess.Popen(['go', 'build'], stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True)
    stdout, stderr = process.communicate()

    # Display the output
    print("Build Output:")
    print(stdout)

    # Display errors, if any
    if stderr:
        print("Build Errors:")
        print(stderr)

    # Check the exit code
    if process.returncode == 0:
        print("Build succeeded.")
    else:
        print("Build failed.")


if __name__ == "__main__":
    main()