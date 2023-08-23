# ECP Encryption/Decryption Demo App
This app demonstrates how to encrypt and decrypt messages using the Enterprise Certificate Proxy (ECP) client library.

# Usage
To use the app, you need to specify the following flags:

mode: The mode of operation, either -e for encryption or -d for decryption.

msg: The message to encrypt or decrypt.

For example, to encrypt the message "Hello, world!", you would run the following command:
    $ go run main.go -mode -e -msg "Hello, world!"
    
The app will print the encrypted message to the console.
To decrypt the message, you would run the following command:
    $ go run main.go -mode -d -msg <encrypted message>
    
The app will print the decrypted message to the console.

# Installation
To properly set up for the demo on a MacOS device, run the following command: 
  $ ./scripts/macos_setup.sh
This script: 
- Creates a self-signed certificate, private key, and MacOS keychain
- Puts the custom keychain on keychain path
- Creates a MacOS ECP config file
