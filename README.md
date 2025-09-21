# 🗄️ FileVault - Cloud Storage Platform

<div align="center">

![FileVault Logo](https://img.shields.io/badge/FileVault-Cloud%20Storage-blue?style=for-the-badge&logo=cloud&logoColor=white)

**A modern, secure cloud storage platform built with cutting-edge technologies**

[![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)](https://golang.org/)
[![React](https://img.shields.io/badge/React-61DAFB?style=flat&logo=react&logoColor=black)](https://reactjs.org/)
[![GraphQL](https://img.shields.io/badge/GraphQL-E10098?style=flat&logo=graphql&logoColor=white)](https://graphql.org/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-336791?style=flat&logo=postgresql&logoColor=white)](https://postgresql.org/)
[![Supabase](https://img.shields.io/badge/Supabase-3ECF8E?style=flat&logo=supabase&logoColor=white)](https://supabase.com/)
[![Render](https://img.shields.io/badge/Render-46E3B7?style=flat&logo=render&logoColor=white)](https://render.com/)
[![Vercel](https://img.shields.io/badge/Vercel-000000?style=flat&logo=vercel&logoColor=white)](https://vercel.com/)

[🚀 Live Demo](https://your-app.vercel.app) • [📖 Documentation](#documentation) • [🛠️ Installation](#installation)

</div>

---

## ✨ Features

### 🔐 **Authentication & Security**
- **JWT-based Authentication** - Secure token-based user sessions
- **Password Protection** - Encrypted user credentials
- **Admin Panel Access** - Password-protected administrative interface
- **User Management** - Complete user registration and login system

### 📁 **File Management**
- **Drag & Drop Upload** - Intuitive file upload with progress tracking
- **Multi-file Upload** - Upload multiple files simultaneously
- **File Organization** - Organized storage with user-specific folders
- **File Preview** - Visual file type indicators with custom icons
- **Download System** - Secure file download with proper authentication
- **File Deletion** - Remove unwanted files with confirmation

### 🎨 **User Interface**
- **Modern Design** - Beautiful, responsive UI with Tailwind CSS
- **Dark/Light Theme** - Elegant gradient backgrounds and glassmorphism effects
- **Interactive Elements** - Smooth animations and hover effects
- **Mobile Responsive** - Optimized for all device sizes
- **Toast Notifications** - Real-time feedback for user actions
- **Search & Filter** - Find files quickly with built-in search

### 👑 **Admin Features**
- **User Overview** - View all registered users and their details
- **File Monitoring** - Complete file listing with uploader information
- **Storage Analytics** - Track storage usage and file statistics
- **User Activity** - Monitor user uploads and file management

### 📊 **Storage Management**
- **Storage Tracking** - Real-time storage usage monitoring
- **File Statistics** - Display total files and formatted storage size
- **Cloud Integration** - Seamless Supabase storage integration
- **Scalable Architecture** - Built to handle growing storage needs

---

## 🏗️ Tech Stack

### **Backend**
- **[Go](https://golang.org/)** - High-performance backend server
- **[GraphQL](https://graphql.org/)** - Flexible API with gqlgen
- **[PostgreSQL](https://postgresql.org/)** - Robust relational database
- **[JWT](https://jwt.io/)** - Secure authentication tokens
- **[Supabase](https://supabase.com/)** - Cloud storage and database services

### **Frontend**
- **[React](https://reactjs.org/)** - Modern UI library with TypeScript
- **[Apollo Client](https://www.apollographql.com/docs/react/)** - GraphQL client with caching
- **[Tailwind CSS](https://tailwindcss.com/)** - Utility-first CSS framework
- **[React Icons](https://react-icons.github.io/react-icons/)** - Beautiful icon library
- **[React Toastify](https://fkhadra.github.io/react-toastify/)** - Elegant notifications
- **[React Router](https://reactrouter.com/)** - Client-side routing

### **Deployment & Infrastructure**
- **[Render](https://render.com/)** - Backend hosting and PostgreSQL database
- **[Vercel](https://vercel.com/)** - Frontend hosting and deployment
- **[Supabase Storage](https://supabase.com/storage)** - File storage and CDN
- **[GitHub](https://github.com/)** - Version control and CI/CD

---

## 🚀 Quick Start

### Prerequisites
- Node.js 18+ and npm
- Go 1.19+
- PostgreSQL
- Git

### Backend Setup

```bash
# Clone the repository
git clone https://github.com/your-username/filevault.git
cd filevault

# Install Go dependencies
go mod download

# Set up environment variables
cp .env.example .env
# Edit .env with your configuration

# Run database migrations
psql "your-database-url" -f setup_db.sql

# Start the server
go run server.go
```

### Frontend Setup

```bash
# Navigate to frontend directory
cd filevault_frontend

# Install dependencies
npm install

# Set up environment variables
cp .env.example .env
# Edit .env with your backend URL

# Start development server
npm start
```

---

## 🔧 Configuration

### Environment Variables

#### Backend (.env)
```env
DATABASE_URL=postgresql://username:password@host:port/database
SUPABASE_URL=https://your-project.supabase.co
SUPABASE_SERVICE_ROLE_KEY=your_service_role_key
JWT_CODE=your_jwt_secret_key
PORT=8080
```

#### Frontend (.env)
```env
REACT_APP_GRAPHQL_ENDPOINT=https://your-backend.onrender.com/query
```

---

## 📱 Screenshots

### 🏠 Dashboard
Beautiful, modern interface with glassmorphism design and intuitive file management.

### 📤 File Upload
Drag and drop interface with real-time progress tracking and multi-file support.

### 👑 Admin Panel
Comprehensive administrative interface for user and file management.

### 📊 Storage Analytics
Real-time storage usage tracking with formatted size display.

---

## 🛠️ API Documentation

### GraphQL Schema

#### Queries
```graphql
type Query {
  users: [User!]!
  userFiles(userID: ID!): [File!]!
  allFiles: [File!]!
  downloadFile(fileID: ID!): File!
  userStorageInfo(userID: ID!): StorageInfo!
}
```

#### Mutations
```graphql
type Mutation {
  createUser(name: String!, email: String!, password: String!): User!
  login(email: String!, password: String!): AuthPayload!
  uploadFile(userID: ID!, file: Upload!): File!
  deleteFile(fileID: ID!): Boolean!
}
```

---

## 🚀 Deployment

### Backend (Render)
1. Connect your GitHub repository to Render
2. Set environment variables in Render dashboard
3. Deploy with automatic builds on push

### Frontend (Vercel)
1. Connect your GitHub repository to Vercel
2. Set `REACT_APP_GRAPHQL_ENDPOINT` environment variable
3. Deploy with automatic builds on push

### Database (Render PostgreSQL)
1. Create PostgreSQL database on Render
2. Run migration scripts to set up tables
3. Connect backend using `DATABASE_URL`

---

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 🙏 Acknowledgments

- **[Supabase](https://supabase.com/)** - For excellent cloud storage and database services
- **[Render](https://render.com/)** - For reliable backend hosting and PostgreSQL
- **[Vercel](https://vercel.com/)** - For seamless frontend deployment
- **[React Icons](https://react-icons.github.io/react-icons/)** - For beautiful, consistent icons
- **[Tailwind CSS](https://tailwindcss.com/)** - For rapid UI development

---

## 📞 Support

If you have any questions or need help, please:

- 📧 Email: support@filevault.com
- 🐛 [Report Issues](https://github.com/your-username/filevault/issues)
- 💬 [Discussions](https://github.com/your-username/filevault/discussions)

---

<div align="center">

**Built with ❤️ by [Your Name](https://github.com/your-username)**

⭐ Star this repository if you found it helpful!

</div>
