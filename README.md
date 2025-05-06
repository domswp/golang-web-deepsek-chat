# 🧠 DeepSeek Web Chat — Go + API + HTML
A lightweight AI chat app built with **Golang**, connected to the **DeepSeek API**, with a minimal web interface.  
Ideal for small projects, assistant demos (e.g., for Westown View apartment), and learning purposes.

---

## 📦 Features
- Chat interface via browser
- Backend in pure Go using OpenAI-compatible SDK
- System prompt: asisten marketing properti (Westown View)
- Customizable prompt, easy to expand
- Ngrok-compatible for easy public access

---

## 📁 Folder Structure

deepseek-chat/
├── main.go # Go backend server
├── .env # API key (DeepSeek)
├── go.mod / go.sum # Go dependencies
└── static/
└── index.html # Simple frontend


---

## ⚙️ Requirements

- Golang (v1.18 or newer)
- Ngrok (optional, for sharing)
- DeepSeek API Key (https://platform.deepseek.com)

---

## 🔧 Installation

```bash
# clone or create folder
mkdir deepseek-chat && cd deepseek-chat
```
# create folders & files
```bash
mkdir static
touch main.go static/index.html .env
```
# init Go module
```bash
go mod init deepseek-chat
```
# install dependencies
```bash
go get github.com/sashabaranov/go-openai
go get github.com/joho/godotenv
go mod tidy
```
## 🧪 .env File

Create .env and add your DeepSeek API key:
```bash
DEEPSEEK_API_KEY=your_deepseek_api_key
```
## 🚀 Run the Project

```bash
go run main.go
```
visit:
```bash
http://localhost:8080
```

## 🌍 Share with Others (Ngrok)
1. Install Ngrok
```bash
wget https://bin.equinox.io/c/bNyj1YtQbY/ngrok-stable-linux-amd64.zip
unzip ngrok-stable-linux-amd64.zip
sudo mv ngrok /usr/local/bin
```
2. Add Auth Token
```bash
ngrok config add-authtoken YOUR_NGROK_TOKEN
```

3. Run Ngrok
```bash
ngrok http 8080
```
Ngrok will give you a public URL like:
```bash
https://1234abcd.ngrok.io
```
Share with your friend

## 🧠 System Prompt

System prompt is defined in main.go:
```bash
{
  Role: "system",
  Content: "Kamu adalah asisten marketing dari Westown View. Tugasmu adalah membantu menjelaskan apartemen Westown View kepada calon customer dengan bahasa yang sopan, informatif, dan meyakinkan.",
}
```
You can easily customize this.

## ✨ Why This Project Was Created

This project was created as a lightweight **web-based AI chatbot** that runs locally but can also be easily shared over the internet using tools like Ngrok — with zero deployment cost.

The main goals are:
- 💼 To demonstrate a practical **AI assistant for real estate marketing**, especially for Westown View Apartment.
- 🧠 To explore the integration of **LLMs like DeepSeek** via OpenAI-compatible APIs using Golang.
- 🌐 To enable **easy sharing of a chatbot** with friends or clients without needing a VPS or cloud server.
- 🎓 To serve as a simple **end-to-end fullstack project**, combining backend API logic and a minimal HTML frontend.

This project is ideal for learning, prototyping, and showcasing how LLMs can be used in everyday real-world applications.
