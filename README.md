# gRPC Payment Plugin

This project implements a gRPC-based payment plugin system with a client and server implementation.

## Project Structure

```
grpc/
├── payment.proto          # Protocol buffer definition
├── main.go               # Client implementation
├── compile.sh            # Script to compile proto files
├── go.mod               # Go module file
├── plugins/
│   └── bankabc/
│       └── main.go      # BankABC plugin server
└── github.com/recitelabs/grpc/plugin/
    ├── payment.pb.go    # Generated message types
    └── payment_grpc.pb.go # Generated gRPC service
```

## Setup

1. **Install dependencies:**
   ```bash
   go mod tidy
   ```

2. **Compile protocol buffers:**
   ```bash
   ./compile.sh
   ```

## Usage

### Running Plugins

You can run different bank plugins using the provided script:

```bash
# Run BankABC plugin (port 50051)
./run_plugin.sh bankabc

# Run BankXYZ plugin (port 50052)
./run_plugin.sh bankxyz

# Run Template plugin (port 50053)
./run_plugin.sh template
```

Or run them directly:

```bash
# BankABC plugin
go run plugins/bankabc/main.go

# BankXYZ plugin  
go run plugins/bankxyz/main.go

# Template plugin
go run plugins/template/main.go
```

### Running the Client

In another terminal, run the client:

```bash
go run main.go
```

By default, the client connects to BankABC (localhost:50051). To connect to a different bank, modify the `bankAddress` variable in `main.go`.

### Testing Multiple Banks

You can test multiple banks by running different plugins on different ports and modifying the client to connect to each one.

## API

The gRPC service provides two methods:

- `SendPayment(PaymentRequest) -> PaymentResponse`
- `CheckStatus(StatusRequest) -> StatusResponse`

### PaymentRequest
- `amount`: Payment amount as string
- `currency`: Currency code (e.g., "ETB")
- `account_number`: Recipient account number
- `recipient_name`: Name of the recipient

### PaymentResponse
- `success`: Boolean indicating if payment was successful
- `transaction_id`: Unique transaction identifier
- `message`: Status message

## Adding New Plugins

To add a new bank plugin, follow these steps:

### 1. Create Plugin Directory

```bash
mkdir -p plugins/your_bank_name
```

### 2. Copy the Template

```bash
cp plugins/template/main.go plugins/your_bank_name/main.go
```

### 3. Customize the Plugin

Edit `plugins/your_bank_name/main.go`:

- Change `BankTemplateServer` to `YourBankServer`
- Update the port number (e.g., `:50054`)
- Add your bank-specific API integration in the `SendPayment` and `CheckStatus` methods
- Update log messages and transaction IDs

### 4. Update the Run Script

Add your plugin to `run_plugin.sh`:

```bash
# Add this case to the script
"your_bank_name")
    PORT=50054
    ;;
```

### 5. Test Your Plugin

```bash
# Run your new plugin
./run_plugin.sh your_bank_name

# Test with the client (update main.go to use your port)
go run main.go
```

### Plugin Structure

Each plugin must implement the `BankPluginServer` interface:

```go
type BankPluginServer interface {
    SendPayment(context.Context, *PaymentRequest) (*PaymentResponse, error)
    CheckStatus(context.Context, *StatusRequest) (*StatusResponse, error)
}
```

## Development

To modify the protocol buffer definition, edit `payment.proto` and run `./compile.sh` to regenerate the Go code.

## Fixed Issues

The following issues were resolved:

1. **Import path errors**: Updated import paths from `"your_project/proto"` to the correct generated package path
2. **Deprecated gRPC calls**: Replaced `grpc.WithInsecure()` with `grpc.WithTransportCredentials(insecure.NewCredentials())`
3. **Missing dependencies**: Added proper Go module initialization and dependencies
4. **Missing context import**: Added context import to the server implementation 
