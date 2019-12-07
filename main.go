package main

import (
	"fmt"
	"strings"
)

var doc = map[string]string{
	"base": `
	\\documentclass[11pt,a4paper,roman]{moderncv}        
	\\moderncvstyle{banking}     
	\\moderncvcolor{blue}               
	\\nopagenumbers{}                                  
	\\usepackage[utf8]{inputenc}
	\\usepackage{fontawesome}
	\\usepackage{tabularx}
	\\usepackage{ragged2e}
	\\usepackage[scale=0.8]{geometry}
	\\usepackage{multicol}
	\\usepackage{import}
	
	
	\\name{@firstName}{@lastName}
	
	\address{@address}{}{}
	  
	\newcommand*{\customcventry}[7][.25em]{
	  \begin{tabular}{@{}l} 
		{\bfseries #4}
	  \end{tabular}
	  \hfill
	  \begin{tabular}{l@{}}
		 {\bfseries #5}
	  \end{tabular} \\
	  \begin{tabular}{@{}l} 
		{\itshape #3}
	  \end{tabular}
	  \hfill
	  \begin{tabular}{l@{}}
		 {\itshape #2}
	  \end{tabular}
	  \ifx&#7&
	  \else{\\
		\begin{minipage}{\maincolumnwidth}
		  \small#7
		\end{minipage}}\fi
	  \par\addvspace{#1}}
	
	\newcommand*{\customcvproject}[4][.25em]{
	  \begin{tabular}{@{}l} 
		{\bfseries #2}
	  \end{tabular}
	  \begin{tabular}{l@{}}
		 {\itshape #3}
	  \end{tabular}
	  \ifx&#4&
	  \else{\\
		\begin{minipage}{\maincolumnwidth}
		  \small#4
		\end{minipage}}\fi
	  \par\addvspace{#1}}
	
	\setlength{\tabcolsep}{12pt}
	
	\begin{document}
	\makecvtitle
	\vspace*{-23mm}
	
	\begin{center}
	\begin{tabular}{ c c c c }
	\faEnvelopeO\enspace @email & \faGithub\enspace suyashdeshpande &  \faMobile\enspace @phone\\  
	\end{tabular}
	\end{center}
		
	\section{EDUCATION}
	
	@educations
	
	\section{KEY SKILLS}
	@skills
	
	\section{ACADEMIC PROJECTS}
	
	@projects
	
	
	\section{EXPERIENCE}
	
	@experiences
		  
	}
	
	\nocite{*}
	\bibliographystyle{plain}
	\bibliography{publications}             
	
	\end{document}
		`,
	"education": `{\customcventry{@date}{@major}{@name}{@location}{}{}}`,
	"project": `
		{\customcvproject{@name}{@startDate - @endDate}
		{\begin{itemize}
		@projectItems
		\end{itemize}
		}
	  }
	  `,
	"item": `
	  \item @itemName
	  `,
	"experience": `
	  {\customcventry{@startDate - @endDate}{@profile}{@name}{@location}{}
	  {\begin{itemize}
		@experienceItems
	  \end{itemize}
	  }
	  }
	  `,
}

//Education represents properties if education
type Education struct {
	Name     string
	Major    string
	Location string
	Date     string
}

//Project represents properties if education
type Project struct {
	Name        string
	StartDate   string
	EndDate     string
	Description []string
}

//Experience represents properties if education
type Experience struct {
	Name        string
	Profile     string
	Location    string
	StartDate   string
	EndDate     string
	Description []string
}

//Form represents properties if education
type Form struct {
	Name       string
	Email      string
	Address    string
	Phone      string
	Educations []Education
	Skills     string
	Projects   []Project
	Experieces []Experience
}

func generate(form Form) string {
	ans := doc["base"]
	firstName := strings.Split(form.Name, " ")[0]
	lastName := strings.Split(form.Name, " ")[1]

	// Replace name in doc
	ans = strings.Replace(ans, "@firstName", firstName, 1)
	ans = strings.Replace(ans, "@lastName", lastName, 1)

	// Replace email in doc
	ans = strings.Replace(ans, "@email", form.Email, 1)

	// Replace address in doc
	ans = strings.Replace(ans, "@address", form.Address, 1)

	// Replace phone in doc
	ans = strings.Replace(ans, "@phone", form.Phone, 1)

	// Replace skills in doc
	ans = strings.Replace(ans, "@skills", form.Skills, 1)

	// Replace education in doc
	educations := ""
	for _, ed := range form.Educations {
		tempEducation := doc["education"]
		tempEducation = strings.Replace(tempEducation, "@name", ed.Name, 1)
		tempEducation = strings.Replace(tempEducation, "@major", ed.Major, 1)
		tempEducation = strings.Replace(tempEducation, "@location", ed.Location, 1)
		tempEducation = strings.Replace(tempEducation, "@date", ed.Date, 1)
		educations += tempEducation
	}
	ans = strings.Replace(ans, "@educations", educations, 1)

	// Replace experiences in doc
	experiences := ""
	for _, exp := range form.Experieces {
		tempExperience := doc["experience"]
		tempExperience = strings.Replace(tempExperience, "@name", exp.Name, 1)
		tempExperience = strings.Replace(tempExperience, "@profile", exp.Profile, 1)
		tempExperience = strings.Replace(tempExperience, "@location", exp.Location, 1)
		tempExperience = strings.Replace(tempExperience, "@startDate", exp.StartDate, 1)
		tempExperience = strings.Replace(tempExperience, "@endDate", exp.EndDate, 1)
		tempD := ""
		for _, d := range exp.Description {
			x := doc["item"]
			x = strings.Replace(x, "@itemName", d, 1)
			tempD += x
		}
		tempExperience = strings.Replace(tempExperience, "@experienceItems", tempD, 1)
		experiences += tempExperience
	}
	ans = strings.Replace(ans, "@experiences", experiences, 1)

	// Replace projects in doc
	projects := ""
	for _, pro := range form.Projects {
		tempProject := doc["project"]
		tempProject = strings.Replace(tempProject, "@name", pro.Name, 1)
		tempProject = strings.Replace(tempProject, "@startDate", pro.StartDate, 1)
		tempProject = strings.Replace(tempProject, "@endDate", pro.EndDate, 1)
		tempD := ""
		for _, d := range pro.Description {
			x := doc["item"]
			x = strings.Replace(x, "@itemName", d, 1)
			tempD += x
		}
		tempProject = strings.Replace(tempProject, "@projectItems", tempD, 1)
		projects += tempProject
	}
	ans = strings.Replace(ans, "@projects", projects, 1)

	return ans
}

func main() {

	temp := Form{
		Name:    "Aniket Singh",
		Address: "Bhopal",
		Email:   "aniketsingh0104@gmail.com",
		Phone:   "9407826614",
		Skills:  "Python | Javascript",
		Educations: []Education{
			Education{
				Name:     "NIT Bhopal",
				Date:     "May 2020",
				Location: "Bhopal",
				Major:    "B.Tech in CSE",
			},
			Education{
				Name:     "School",
				Date:     "April 2016",
				Location: "Bhopal",
				Major:    "12th Class",
			},
		},
		Projects: []Project{
			Project{
				Name:        "Latex Resume Generator",
				StartDate:   "1 Dec 2019",
				EndDate:     "15 Dec 2019",
				Description: []string{"It is a resume generator"},
			},
		},
		Experieces: []Experience{
			Experience{
				Name:        "JP Morgan Chase",
				Profile:     "SDE Intern",
				Location:    "Bombay",
				StartDate:   "1 May 2019",
				EndDate:     "28 June 2019",
				Description: []string{"Developed a chatbot in python"},
			},
		},
	}
	fmt.Println("Hello World")
	fmt.Println(generate(temp))
}
