const axios = require('axios');
const fs = require('fs').promises;

const API_URL = 'http://localhost:8080/encrypt';

async function encryptAndStoreData() {
    try {
        // Replace with actual data to encrypt
        const rawData = "Your data to encrypt";

        // Send data to the API for encryption
        const response = await axios.post(API_URL, { data: rawData });
        const { encryptedData, key, nonce } = response.data;

        // Store the encrypted data
        await fs.writeFile('encryptedData.bin', encryptedData, 'binary');

        // Store the keys
        const keys = { key: key.toString('hex'), nonce: nonce.toString('hex') };
        await fs.writeFile('encryptionKeys.json', JSON.stringify(keys), 'utf8');

        console.log('Encrypted data and keys stored successfully.');
    } catch (error) {
        console.error('Error:', error.message);
    }
}

encryptAndStoreData();


encryptAndStoreData();