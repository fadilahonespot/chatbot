# Gunakan golang:alpine sebagai base image
FROM golang:alpine

# Set working directory
WORKDIR /app

# Membersihkan cache modul Go di dalam kontainer.
RUN go clean --modcache

# Copy kode Go Anda ke direktori /app di dalam kontainer
COPY . .

# Build aplikasi Go
RUN go build -o my-go-app

# Port yang akan digunakan oleh aplikasi
EXPOSE 8086

# Perintah untuk menjalankan aplikasi saat kontainer dijalankan
CMD ["./my-go-app"]