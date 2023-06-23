# Learnado: Igniting Minds, Inspiring Learning (Content Manager Edition)

## Overview

Named in honor of the genius inventor, Leonardo da Vinci, Learnado is an inventive open-source educational platform that fosters a love of learning in every corner of the globe. Merging the worlds of creativity and education, Learnado is designed to empower learners and educators alike, transcending traditional boundaries of education. Learnado is built with a strong focus on accessibility, enabling learning on a multitude of devices and operating systems. It is built for offline usage, ensuring that education is not limited by internet availability. This makes Learnado an ideal choice for those living in areas with limited internet access or in situations where internet connectivity is unreliable. This project aims to democratize education and provide a quality learning experience to everyone, everywhere.

## Getting Started

As a content manager, you will be responsible for creating and managing the course content. Here are the steps to get started:

### 1. Downloading Learnado

Before you can start using Learnado, you need to download the Learnado software. Follow these steps to download Learnado from the GitHub releases:

1. Navigate to the "Releases" section of the repository.
2. Locate the latest stable release version of Learnado.
3. Download the Learnado release package specific to your operating system (e.g., macOS, Linux, or Windows).
4. Extract the downloaded package to a location on your computer where you want to store the Learnado files.

Once you have downloaded and extracted the Learnado package, you can proceed with the remaining steps.

### 2. Install Hugo

Before you start structuring and customizing your course content, you need to install Hugo, the static site generator that powers Learnado's course rendering. Follow these steps to install Hugo:

1. Visit the official Hugo website at [https://gohugo.io](https://gohugo.io).
2. Navigate to the "Get Started" section and choose the installation method that is compatible with your operating system.
3. Follow the provided installation instructions for your specific operating system.
4. Verify the successful installation of Hugo by opening a terminal or command prompt and running the command `hugo version`. If Hugo is properly installed, it will display the installed version.

Once Hugo is installed on your system, you can proceed with the course structuring and customization steps mentioned in the following sections.

Note: Familiarity with basic command-line operations and Markdown syntax will be helpful when working with Hugo and structuring your course content.

### 3. Setting Up the Home Page

To create an engaging landing page, you'll need to set up the `homepage.md` file. The `homepage.md` file serves as the main page of Learnado.

Remember to refer to the Hugo Theme Learn documentation for additional guidance on structuring your home page and utilizing the theme's features to create an engaging learning experience.

### 4. Course Creation

Learnado utilizes the powerful Hugo Theme Learn to structure and present the course content. 

Courses are created by instructors writing Markdown files. These files form the basis of your course content. If you're new to Markdown, don't worry! It's a simple and straightforward language designed for easy readability and writing.

To ensure optimal course organization and customization, content managers should refer to the official Hugo Theme Learn documentation available at [https://github.com/matcornic/hugo-theme-learn](https://github.com/matcornic/hugo-theme-learn).

The Hugo Theme Learn website provides detailed instructions on how to structure your course materials using the theme. It covers topics such as creating sections, adding lessons, configuring navigation, and customizing the visual appearance of your courses.

To align the Learnado platform with your specific needs and branding, you can customize the "hugo" folder provided by Learnado. This folder contains the necessary files and configurations related to the Hugo Theme Learn. Feel free to modify this folder according to your requirements, including changing the theme's colors, fonts, layouts, and more.

By following the guidelines and customizing the Learnado "hugo" folder, you can create a unique and engaging learning experience for your students.

Note: It's important to familiarize yourself with the Hugo framework and the Hugo Theme Learn documentation to make the most of the customization options available.

### 5. Executing the Learnado Binary

To run Learnado on your system, you need to execute the Learnado binary. Here are the steps to execute the binary on different operating systems:

#### For MacOS and Linux:
1. Open a terminal.
2. Navigate to the directory where the Learnado binary is located.
3. Use the `cd` command followed by the directory path to navigate to the Learnado folder. For example:
   ```
   cd /path/to/learnado
   ```
4. Once inside the Learnado folder, run the Learnado binary by executing the following command:
   ```
   ./Learnado-ContentManager
   ```
   Note: If you encounter permission issues, you may need to make the Learnado binary executable by running the command `chmod +x Learnado-ContentManager` before executing it.

#### For Windows:
1. Open the command prompt or PowerShell.
2. Navigate to the directory where the Learnado binary is located. This may be the folder where you downloaded or cloned the Learnado repository.
3. Use the `cd` command followed by the directory path to navigate to the Learnado folder. For example:
   ```
   cd C:\path\to\learnado
   ```
4. Once inside the Learnado folder, run the Learnado binary by executing the following command:
   ```
   Learnado-ContentManager.exe
   ```

After executing the Learnado binary, the Learnado platform will start running on your local machine. If it successfully runs, you will see this message: `⇨ http server started on [::]:8080`.

### 6. Course Upload and Registration

Once a course has been created, it needs to be uploaded to the server that is running Learnado. After uploading it, open a browser and go to [http://localhost:8080](http://localhost:8080) (if you are using the default port) to view the Learnado Content Manager Graphical User Interface (GUI). Click on the "Course Management" tab and input the necessary details such as the course name and folder path in the fields that are above the "Create Course" button. With these details, the course is registered on the Learnado platform and is ready to be distributed to students.

### 7. License Generation and Distribution

After a course is registered, licenses for the course can be generated with the same GUI. These licenses let Learnado know which students are authorized to download and access the course. Once the licenses are distributed and activated by the students, the course will be downloaded to their devices.

## Features

### Run Anywhere

Learnado is designed to run on a wide variety of devices and operating systems, including Windows, Linux, and MacOS. This ensures that as many students as possible can access the learning materials, regardless of their device or operating system.

### Offline Learning

All course materials are packaged for offline use. Once a course is downloaded, Learnado works entirely offline, making it perfect for areas with limited or unreliable internet access.

### Content Protection

Course content is compressed and encrypted before being sent to the student's software. This protects against piracy and unauthorized access, ensuring that your content is safe and secure.

### Periodic Updates

Learnado periodically checks for new course content, ensuring that students always have access to the most up-to-date materials. The update checks require an internet connection, but once the updates are downloaded, they can be accessed offline.

Thank you for choosing Learnado as your educational platform. Together, we can bring quality education to everyone, everywhere!