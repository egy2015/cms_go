## Programming Manual: CMS_GO

This manual guides you through setting up the development environment for your Go project named "[Your Project Name]".

**Prerequisites:**

* Go installed on your system. You can download and install Go from the official website: [https://go.dev/doc/install](https://go.dev/doc/install)

**Steps:**

1. **Initialize Go Module:**

   Open a terminal in the directory where you want to create your project. Run the following command, replacing `[Your Project Name]` with your actual project name:

   ```bash
   go mod init [Your Project Name]
   ```

   This command creates a `go.mod` file in your project directory, which manages dependencies for your project.

2. **Generate Folder Structure:**

   Run the following command to generate a basic folder structure for your project:

   ```bash
   make folder_init
   ```

3. **Happy Coding:**

   You can use your own folder structure / base, and you can type `make` and you can see what commands are available to speed up your progress.

**Additional Notes:**

* It's better if you use version control for any programming language, well in this case `Go Language`.
* If you want to use Golang Version Control, i recommend using this `https://github.com/moovweb/gvm`

**Customization:**

* Remember to replace `[Your Project Name]` with your actual project name throughout this process.
* You can further customize the generated code and folder structure based on your project's specific needs.

This manual provides a basic overview of setting up your Go development environment. As you progress with your project, you'll encounter more specific commands and procedures related to programming in Go.
