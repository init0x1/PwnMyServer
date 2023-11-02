# PwnMyServer

a tool that allows you to execute commands on a remote server and retrieve the results as JSON data.

## Installation

To install the PwnMyServer , follow these steps:

### Clone this GitHub repository to your local machine:

```
git clone https://github.com/0xAbdoAli/PwnMyServer.git
```

### Run the tool:

```
go run main.go
```

## Usage

The PwnMyServer provides a simple interface for sending commands to a remote server and receiving JSON responses. Here's an example :

```
Enter command: uname -a
```

### output like :

```json 
{
    "command-result": [
        "Linux srv-ckvrs3qfh5gs73e5k1d0-hibernate-77cdbbcf5c-j96wd 6.2.0-1014-aws #14~22.04.1-Ubuntu SMP Thu Oct  5 22:43:45 UTC 2023 x86_64 GNU/Linux",
        ""
    ]
}
```
