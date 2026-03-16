Groupie Tracker

A dynamic web application built with Go that consumes a multi-part RESTful API to visualize data about musical artists, their concert locations, and scheduled dates. This project demonstrates full-stack development using only the Go standard library, focusing on data orchestration, JSON manipulation, and client-server interactions.
🚀 Live Demo

Explore the Groupie Tracker Live

🛠️ Features

    Data Orchestration: Successfully fetches and unmarshals four distinct API endpoints (Artists, Locations, Dates, and Relations).

    Interactive Visualization:

        Artist Gallery: A clean, responsive grid of clickable cards displaying artist images and names.

        Detail Views: Dedicated pages for each artist showcasing their history, members, and a consolidated view of concert relations.

    Search Functionality: A real-time filtered search bar to help users find specific artists quickly.

    Robust Backend: Built entirely in Go, handling concurrent data fetching and structured error management to ensure zero server crashes.

    Responsive UI: Styled with a "Midnight Blue" and "Cyan" theme, optimized for both desktop and mobile viewing.

🏗️ Technical Stack

    Backend: Go (Golang) — Standard Library only

    Frontend: HTML5, CSS3, Go Templates (html/template)

    Data Format: JSON

    Deployment: Render / GitHub

📋 Project Structure
Plaintext

groupie-tracker/
├── main.go              # Server entry point & routing
├── backend/             # Core Logic & Data Management
│   ├── handlers.go      # HTTP handlers for Home and Details
│   ├── models.go        # Struct definitions for JSON mapping
│   └── services.go      # API fetching and JSON unmarshalling
├── frontend/            # UI Layer
│   ├── index.html       # Homepage template
│   ├── details.html     # Artist details template
│   ├── error.html       # Error handling template
│   └── static/          # Assets
│       └── styles.css   # Custom CSS styling
├── go.mod               # Dependency management
└── README.md            # Documentation

💻 Installation & Usage

To run this project locally, ensure you have Go installed.

    Clone the repository:
    Bash

    git clone https://github.com/eojuma/groupie-tracker.git
    cd groupie-tracker

    Run the application:
    Bash

    go run main.go

    Access the site:
    Open your browser and navigate to http://localhost:8080

🎓 Learning Objectives

This project was developed as part of the Zone01 curriculum to master:

    JSON Manipulation: Handling complex nested JSON structures and mapping them to Go structs.

    Client-Server Architecture: Implementing the Request-Response cycle and managing state between the backend and the browser.

    Event Handling: Triggering server-side actions via client-side interactions (e.g., clicking an artist card to fetch relational data).

👥 The Team

    Evans Juma - Lead / UI Design

    Claire Gisore - Backend Contributor

    Ibrahim Samwel - Backend Contributor
