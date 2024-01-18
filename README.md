# Chatbot API (Golang)

## Introduction
This repository contains the source code for a Chatbot API based on Golang. The API utilizes a chatbot powered by OpenAI and is designed to provide a conversational interface.

## Setup
To run the API, you need to set up the environment variables. Create a file named `.env` based on the provided `env.example` file. Update the values with your specific configurations.

1. Running in local computer
    Make sure to replace placeholders such as your_db_username, your_db_password, etc., with your actual configurations.
    ```env
    # Application
    APP_PORT=5067

    # Database
    DB_USERNAME=your_db_username
    DB_PASSWORD=your_db_password
    DB_HOST=your_db_host
    DB_PORT=3306
    DB_NAME=your_db_name
    DB_DEBUG=false

    # Logger
    LOGGER_LOGS_WRITE=true
    LOGGER_FOLDER_PATH="./logs"

    # Redis
    REDIS_HOST=your_redis_host
    REDIS_PORT=6379
    REDIS_PASSWORD=your_redis_password
    REDIS_USERNAME=default

    # OpenAI
    OPEN_AI_TOKEN=your_openai_token

    ```
    After setting up the environment variables, you can run the API using the following commands:
    ```
    go run main.go
    ```

2. Running in docker container
    Make sure to fill in the placholders below with the values you like (except for OPEN_AI_TOKEN, you must fill it with your openAi token) and delete the rest of the env that is not listed below.
    ```env
    DB_USERNAME=db_username
    DB_PASSWORD=db_password
    DB_NAME=db_name

    REDIS_PASSWORD=redis_password

    OPEN_AI_TOKEN=your_openai_token
    ```
    After setting up the environment variables, use the following command to build and run your application in Docker:
    - use make command
    ```
    make build
    ```

    - use docker command
    ```
    docker-compose up
    ```

## Endpoints

1. Register
    - Request
        - Method: POST
        - URL: localhost:5067/register
        - Body:
        ```json
        {
            "name": "fadilah",
            "email": "fadilah65@gmail.com",
            "password": "123456"
        }
        ```

    - Response
        - Status: OK (200)
        - Body:
        ```json
        {
            "code": 200,
            "message": "Success",
            "data": null
        }
        ```

2. Login
    - Request
        - Method: POST
        - URL: localhost:5067/login
        - Body:
        ```json
        {
            "email": "fadilah65@gmail.com",
            "password": "123456"
        }

        ```

    - Response
        - Status: OK (200)
        - Body:
        ```json
        {
            "code": 200,
            "message": "Success",
            "data": {
                "id": 5,
                "name": "fadilah",
                "email": "fadilah620@gmail.com",
                "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDU1NDIxMzksInVzZXJJZCI6NX0.wl24af17wQy6e8HDX4qW4fRO6SIsQozPWMIdbwFtfqM"
            }
        }
        ```

3. Chat Question
    - Request
        - Method: POST
        - URL: localhost:5067/chat
        - Headers:
            - Authorization: Bearer {{access-token}}
        - Body:
        ```json
        {
            "question": "tools yang di butuhkan untuk koding?"
        }
        ```

    - Response
        - Status: OK (200)
        - Body:
        ```json
        {
            "code": 200,
            "message": "Success",
            "data": {
                "answer": "Berikut ini adalah langkah-langkah umum untuk mencuci mobil: ... (detailed answer)"
            }
        }
        ```

4. Chat History
    - Request
        - Method: GET
        - URL: localhost:5067/chat
        - Headers:
            - Authorization: Bearer {{access-token}}
    - Response
        - Status: OK (200)
        - Body:
        ```json
        {
            "code": 200,
            "message": "Success",
            "data": [
                {
                    "id": 2,
                    "name": "Bot",
                    "message": "Untuk koding, ada beberapa tools yang umumnya digunakan oleh para pengembang. ... (detailed answer)"
                },
                {
                    "id": 1,
                    "name": "fadilah",
                    "message": "tools yang di butuhkan untuk koding?"
                }
            ]
        }
        ```
    
    ## Unit Testing
    To run unit tests, use the following command:
    ```
    make test
    ```

    ## Sample Logs
    - Sys Logs
    ```
    {"time":"2024-01-18T10:11:41.157239+07:00","level":"INFO","msg":"Incoming Request","SYS":{"app_name":"chatbot-service","app_version":"1.0.0","app_port":5067,"app_thread_id":"a1de5e81-f68a-4a89-adeb-ad1637060a50","header":{"Accept":["*/*"],"Accept-Encoding":["gzip, deflate, br"],"Authorization":["Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDU2MzExNTYsInVzZXJJZCI6MX0.UzkRwnunGSlC4Lwy05bA5JWaHIxXBWOA5zEfGNlNiXo"],"Connection":["keep-alive"],"Content-Length":["58"],"Content-Type":["application/json"],"Postman-Token":["43eb6f1e-df07-4c8a-9107-a1308800ccf9"],"User-Agent":["PostmanRuntime/7.36.1"]},"app_method":"POST","app_uri":"/chat"}}
    {"time":"2024-01-18T10:11:41.159076+07:00","level":"INFO","msg":"[REQUEST]","SYS":{"app_name":"chatbot-service","app_version":"1.0.0","app_port":5067,"app_thread_id":"a1de5e81-f68a-4a89-adeb-ad1637060a50","header":{"Accept":["*/*"],"Accept-Encoding":["gzip, deflate, br"],"Authorization":["Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDU2MzExNTYsInVzZXJJZCI6MX0.UzkRwnunGSlC4Lwy05bA5JWaHIxXBWOA5zEfGNlNiXo"],"Connection":["keep-alive"],"Content-Length":["58"],"Content-Type":["application/json"],"Postman-Token":["43eb6f1e-df07-4c8a-9107-a1308800ccf9"],"User-Agent":["PostmanRuntime/7.36.1"]},"app_method":"POST","app_uri":"/chat"},"atribute":{"message_0":{"question":"tools yang di butuhkan untuk koding?"}}}
    {"time":"2024-01-18T10:11:41.190441+07:00","level":"INFO","msg":"GeneratText REQUEST","SYS":{"app_name":"chatbot-service","app_version":"1.0.0","app_port":5067,"app_thread_id":"a1de5e81-f68a-4a89-adeb-ad1637060a50","header":{"Accept":["*/*"],"Accept-Encoding":["gzip, deflate, br"],"Authorization":["Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDU2MzExNTYsInVzZXJJZCI6MX0.UzkRwnunGSlC4Lwy05bA5JWaHIxXBWOA5zEfGNlNiXo"],"Connection":["keep-alive"],"Content-Length":["58"],"Content-Type":["application/json"],"Postman-Token":["43eb6f1e-df07-4c8a-9107-a1308800ccf9"],"User-Agent":["PostmanRuntime/7.36.1"]},"app_method":"POST","app_uri":"/chat"},"atribute":{"message_0":{"model":"gpt-3.5-turbo","messages":[{"role":"system","content":"Halo! Perkenalkan aku adalah ChatBot Assistant. Bagimana aku bisa membantumu hari ini?"},{"role":"user","content":"tools yang di butuhkan untuk koding?"},{"role":"user","content":"tools yang di butuhkan untuk koding?"}]}}}
    {"time":"2024-01-18T10:11:56.17777+07:00","level":"INFO","msg":"GenerateText RESPONSE","SYS":{"app_name":"chatbot-service","app_version":"1.0.0","app_port":5067,"app_thread_id":"a1de5e81-f68a-4a89-adeb-ad1637060a50","header":{"Accept":["*/*"],"Accept-Encoding":["gzip, deflate, br"],"Authorization":["Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDU2MzExNTYsInVzZXJJZCI6MX0.UzkRwnunGSlC4Lwy05bA5JWaHIxXBWOA5zEfGNlNiXo"],"Connection":["keep-alive"],"Content-Length":["58"],"Content-Type":["application/json"],"Postman-Token":["43eb6f1e-df07-4c8a-9107-a1308800ccf9"],"User-Agent":["PostmanRuntime/7.36.1"]},"app_method":"POST","app_uri":"/chat"},"atribute":{"message_0":{"id":"chatcmpl-8iD6czURq3p7ynfwTXG8IHvl1nZ8N","object":"chat.completion","created":1705547502,"model":"gpt-3.5-turbo-0613","choices":[{"index":0,"message":{"role":"assistant","content":"Untuk memulai koding, ada beberapa tools yang biasanya digunakan oleh para pengembang. Berikut beberapa tools yang biasa digunakan:\n\n1. Text Editor atau Integrated Development Environment (IDE): seperti Visual Studio Code, Sublime Text, Atom, atau IntelliJ IDEA. Tools ini digunakan untuk menulis dan mengedit kode.\n\n2. Bahasa Pemrograman: Pilihlah bahasa pemrograman yang ingin kamu pelajari atau gunakan. Contohnya, Python, JavaScript, Java, atau PHP.\n\n3. Command Line Interface (CLI): Untuk menjalankan perintah atau skrip dari baris perintah, seperti Command Prompt di Windows atau Terminal di macOS dan Linux.\n\n4. Version Control System (VCS): Berguna untuk mengatur versi dan kolaborasi dengan tim pengembang lain. Git adalah salah satu VCS yang populer.\n\n5. Browser: Untuk menguji dan mengembangkan aplikasi web, kamu memerlukan browser seperti Google Chrome atau Mozilla Firefox.\n\n6. Dokumentasi: Selalu periksa dokumentasi resmi bahasa pemrograman atau framework yang kamu gunakan, seperti dokumentasi Python atau dokumentasi ReactJS.\n\n7. Stack Overflow dan Forum Diskusi: Bergabung dalam komunitas pengembang dan bergabunglah dalam forum diskusi seperti Stack Overflow untuk mencari jawaban atas pertanyaan atau masalah yang kamu hadapi.\n\nItulah beberapa tools dasar yang sering digunakan dalam proses pengembangan aplikasi. Semoga membantu!"},"finish_reason":"stop"}],"usage":{"prompt_tokens":59,"completion_tokens":326,"total_tokens":385},"system_fingerprint":""}}}
    {"time":"2024-01-18T10:11:56.231086+07:00","level":"INFO","msg":"[RESPONSE]","SYS":{"app_name":"chatbot-service","app_version":"1.0.0","app_port":5067,"app_thread_id":"a1de5e81-f68a-4a89-adeb-ad1637060a50","header":{"Accept":["*/*"],"Accept-Encoding":["gzip, deflate, br"],"Authorization":["Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDU2MzExNTYsInVzZXJJZCI6MX0.UzkRwnunGSlC4Lwy05bA5JWaHIxXBWOA5zEfGNlNiXo"],"Connection":["keep-alive"],"Content-Length":["58"],"Content-Type":["application/json"],"Postman-Token":["43eb6f1e-df07-4c8a-9107-a1308800ccf9"],"User-Agent":["PostmanRuntime/7.36.1"]},"app_method":"POST","app_uri":"/chat"},"resp":{"code":200,"data":{"answer":"Untuk memulai koding, ada beberapa tools yang biasanya digunakan oleh para pengembang. Berikut beberapa tools yang biasa digunakan:\n\n1. Text Editor atau Integrated Development Environment (IDE): seperti Visual Studio Code, Sublime Text, Atom, atau IntelliJ IDEA. Tools ini digunakan untuk menulis dan mengedit kode.\n\n2. Bahasa Pemrograman: Pilihlah bahasa pemrograman yang ingin kamu pelajari atau gunakan. Contohnya, Python, JavaScript, Java, atau PHP.\n\n3. Command Line Interface (CLI): Untuk menjalankan perintah atau skrip dari baris perintah, seperti Command Prompt di Windows atau Terminal di macOS dan Linux.\n\n4. Version Control System (VCS): Berguna untuk mengatur versi dan kolaborasi dengan tim pengembang lain. Git adalah salah satu VCS yang populer.\n\n5. Browser: Untuk menguji dan mengembangkan aplikasi web, kamu memerlukan browser seperti Google Chrome atau Mozilla Firefox.\n\n6. Dokumentasi: Selalu periksa dokumentasi resmi bahasa pemrograman atau framework yang kamu gunakan, seperti dokumentasi Python atau dokumentasi ReactJS.\n\n7. Stack Overflow dan Forum Diskusi: Bergabung dalam komunitas pengembang dan bergabunglah dalam forum diskusi seperti Stack Overflow untuk mencari jawaban atas pertanyaan atau masalah yang kamu hadapi.\n\nItulah beberapa tools dasar yang sering digunakan dalam proses pengembangan aplikasi. Semoga membantu!"},"message":"Success"}}

    ```

    - TDR logs
    ```
    {"time":"2024-01-18T10:11:56.231262+07:00","level":"INFO","msg":"TDR","TDR":{"request_id":"a1de5e81-f68a-4a89-adeb-ad1637060a50","path":"/chat","method":"POST","port":5067,"rt":15074,"rc":"200","header":{"Accept":"*/*","Accept-Encoding":"gzip, deflate, br","Authorization":"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDU2MzExNTYsInVzZXJJZCI6MX0.UzkRwnunGSlC4Lwy05bA5JWaHIxXBWOA5zEfGNlNiXo","Connection":"keep-alive","Content-Length":"58","Content-Type":"application/json","Postman-Token":"43eb6f1e-df07-4c8a-9107-a1308800ccf9","User-Agent":"PostmanRuntime/7.36.1"},"req":{"question":"tools yang di butuhkan untuk koding?"},"resp":{"code":200,"data":{"answer":"Untuk memulai koding, ada beberapa tools yang biasanya digunakan oleh para pengembang. Berikut beberapa tools yang biasa digunakan:\n\n1. Text Editor atau Integrated Development Environment (IDE): seperti Visual Studio Code, Sublime Text, Atom, atau IntelliJ IDEA. Tools ini digunakan untuk menulis dan mengedit kode.\n\n2. Bahasa Pemrograman: Pilihlah bahasa pemrograman yang ingin kamu pelajari atau gunakan. Contohnya, Python, JavaScript, Java, atau PHP.\n\n3. Command Line Interface (CLI): Untuk menjalankan perintah atau skrip dari baris perintah, seperti Command Prompt di Windows atau Terminal di macOS dan Linux.\n\n4. Version Control System (VCS): Berguna untuk mengatur versi dan kolaborasi dengan tim pengembang lain. Git adalah salah satu VCS yang populer.\n\n5. Browser: Untuk menguji dan mengembangkan aplikasi web, kamu memerlukan browser seperti Google Chrome atau Mozilla Firefox.\n\n6. Dokumentasi: Selalu periksa dokumentasi resmi bahasa pemrograman atau framework yang kamu gunakan, seperti dokumentasi Python atau dokumentasi ReactJS.\n\n7. Stack Overflow dan Forum Diskusi: Bergabung dalam komunitas pengembang dan bergabunglah dalam forum diskusi seperti Stack Overflow untuk mencari jawaban atas pertanyaan atau masalah yang kamu hadapi.\n\nItulah beberapa tools dasar yang sering digunakan dalam proses pengembangan aplikasi. Semoga membantu!"},"message":"Success"},"error":""}}
    ```
