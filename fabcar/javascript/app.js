/*
 * Copyright IBM Corp. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

'use strict';

const { Gateway, Wallets } = require('fabric-network');
const fs = require('fs');
const path = require('path');
async function main() {
    try {
        // load the network configuration
        const ccpPath = path.resolve(__dirname, '..', '..', 'test-network', 'organizations', 'peerOrganizations', 'org1.example.com', 'connection-org1.json');
        let ccp = JSON.parse(fs.readFileSync(ccpPath, 'utf8'));

        // Create a new file system based wallet for managing identities.
        const walletPath = path.join(process.cwd(), 'wallet');
        const wallet = await Wallets.newFileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);

        // Check to see if we've already enrolled the user.
        const identity = await wallet.get('appUser');
        if (!identity) {
            console.log('An identity for the user "appUser" does not exist in the wallet');
            console.log('Run the registerUser.js application before retrying');
            return;
        }

        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();
        await gateway.connect(ccp, { wallet, identity: 'appUser', discovery: { enabled: true, asLocalhost: true } });

        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork('mychannel');

        // Get the contract from the network.
        const contract = network.getContract('fabcar');

        // Submit the specified transaction.
        // createCar transaction - requires 5 argument, ex: ('createCar', 'CAR12', 'Honda', 'Accord', 'Black', 'Tom')
        // changeCarOwner transaction - requires 2 args , ex: ('changeCarOwner', 'CAR12', 'Dave')
        
        // await contract.submitTransaction('StoreCode', 'www.sp.sust.com', 'www.idp.sust.com','1234567');
        // console.log('Transaction has been submitted for store code1');
        // await contract.submitTransaction('StoreCode', 'www.sp1.sust.com', 'www.idp3.sust.com','1232567');
        // console.log('Transaction has been submitted for store code2');
        // await contract.submitTransaction('StoreCode', 'www.sp2.sust.com', 'www.idp.sust.com','1234967');
        // console.log('Transaction has been submitted for store code3');
        // console.log(`\\\\\\\\\\\\\\\\\n`);

         await contract.submitTransaction('StoreMetaData', 'www.idp.sust.com','eshita');
        // console.log('Transaction has been submitted');
        // await contract.submitTransaction('StoreMetaData', 'www.idp3.sust.com','honda');
        // console.log('Transaction has been submitted metaDAta 3');
        // await contract.submitTransaction('StoreMetaData', 'www.idp4.sust.com','honda');
        // console.log('Transaction has been submitted metaData 5');
        // await contract.submitTransaction('StoreMetaData', 'www.idp5.sust.com','honda');
        // console.log('Transaction has been submitted metaData 6');
    
        console.log(`\\\\\\\\\\\\\\\\\n`);
     //    await contract.submitTransaction('CreateCar', 'CAR12', 'Honda', 'Accord', 'Black', 'Tom');
        // console.log('Transaction has been submitted');
        // const result = await contract.evaluateTransaction('QueryAllCars');
        // console.log(`Transaction has been evaluated, result is: ${result.toString()}`);
        console.log(`\\\\\\\\\\\\\\\\\n`);
        // const result1 = await contract.evaluateTransaction('AllMetaData');
        // console.log(`Transaction has been evaluated for all Metadata, result is: ${result1.toString()}\n`);
        // console.log(`\\\\\\\\\\\\\\\\\n`);
        // await contract.submitTransaction('ChangeMetaData', 'www.idp.sust.com', 'Eshita');
        // console.log('Transaction has been submitted');

        console.log(`\\\\\\\\\\\\\\\\\n`);

       // const result15 = await contract.evaluateTransaction('QueryCode', 'www.idp3.sust.com');
      //  console.log(`Transaction has been evaluated, result is: ${result15.toString()}\n`);2
        // const result16 = await contract.evaluateTransaction('QueryMetaData', 'www.idp.sust.com');
        // console.log(`Transaction has been evaluated, result16 is: ${result16.toString()}\n`);

        // const result11 = await contract.evaluateTransaction('AllCodeData');
        //  console.log(`Transaction has been evaluated, result is: ${result11.toString()}\n`);


        //  const result112 = await contract.evaluateTransaction('CodeFetch', 'www.sp1.sust.com', 'www.idp3.sust.com');
        //  console.log(`Transaction has been evaluated, result is: ${result112.toString()}\n`);

        await contract.submitTransaction('StoreTalList', 'www.sp2.sust.com', 'www.idp3.sust.com');
         await contract.submitTransaction('StoreTalList', 'www.sp2.sust.com', 'www.idp2.sust.com');
        
     //    const resultq = await contract.evaluateTransaction('TalListFetch', 'www.sp1.sust.com');
        
        // console.log(`Transaction has been evaluated, result is TalList: ${ resultq.values()}\n`);
        
         const resultop = await contract.evaluateTransaction('TalListReturn', 'www.sp2.sust.com');
         let resultJson = JSON.parse(resultop.toString())
         let arr = []
         for (var i in resultJson)
            {
                var name = resultJson[i].Tal;
                arr.push(name)
            }
         console.log(arr)
        //  console.log(`Transaction has been evaluated, result is TalList: ${v[0].Tal}\n`);
         
        //  for (let [key, value] of Object.entries(resultop.toString())) {
        //     console.log(key, value);
        // }
      //   await contract.submitTransaction('StoreTalList', 'www.sp1.sust.com', 'www.idp2.sust.com');
           
        //  const resultp = await contract.evaluateTransaction('TalListReturn', 'www.sp2.sust.com');
        //  console.log(`Transaction has been evaluated, result is TalList: ${resultp.toString().valueOf(tal)}\n`);

        // await contract.submitTransaction('DeleteCodeSp', 'www.sp1.sust.com', 'www.idp3.sust.com');
        //  const result113 = await contract.evaluateTransaction('AllCodeData');
        //  console.log(`Transaction has been evaluated, result is: ${result113.toString()}\n`);

        // await contract.submitTransaction('ChangeCarOwner','CAR12','eshita');
    ///   console.log(`Transaction has been evaluated, result is: ${result3.toString()}\n`);

    // const result12 = await contract.evaluateTransaction('QueryForSpecificUser', 'www.idp.sust.com' );
    // console.log(`Transaction has been evaluated for specific user www.idp.sust.com  , result is: ${result12.toString()}\n`);
        
    // const result09 = await contract.evaluateTransaction('UserFetch', 'www.idp.sust.com' );
    // console.log(`Transaction has been evaluated for user fetch www.idp.sust.com  , result is: ${result09.toString()}\n`);
    
    const result90 = await contract.evaluateTransaction('MetaDataFetch', 'www.idp.sust.com' );
    console.log(`Transaction has been evaluated for metadataFetch www.idp.sust.com  , result is: ${result90.toString()}\n`);
    

    // const result10 = await contract.evaluateTransaction('MetaDataFetch', 'www.sp1.sust.com' );
    // console.log(`Transaction has been evaluated for , result is: ${result10.toString()}\n`);

        // Disconnect from the gateway.
        await gateway.disconnect();

    } catch (error) {
        console.error(`Failed to submit transaction: ${error}`);
        process.exit(1);
    }
}

main();
