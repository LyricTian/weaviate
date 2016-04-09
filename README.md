# Weaviate [![Build Status](https://travis-ci.org/weaviate/weaviate.svg)](https://travis-ci.org/weaviate/weaviate)
**A private cloud for the Internet of Things [Weave](https://developers.google.com/weave) protocol**

Weaviate is a REST API based software-as-a-service platform that is able to process the Google Weave protocol. It can be used in combination with all Google's Weave and Brillo libraries ([link to the libs](https://weave.googlesource.com/), [link to Weave](https://developers.google.com/weave), [link to Brillo](https://developers.google.com/brillo)).

You can use Weaviate on simple local machines, or complex distributed networks with Node in combination with a Cassandra database.

*Note:
Weaviate is not fully testable / production ready yet. You can follow the progress for the first release candidates [here](https://github.com/weaviate/weaviate/milestones). Follow this repo or sign up for the [mailing list](http://eepurl.com/bRsMir) to stay informed about the progress.*

![NPM](https://nodei.co/npm/weaviate.png?downloads=true "NPM")

### Table of contents
* [How does it work?](#how-does-it-work)
* [Installation](#installation)
* [Using the weaviate() function](#using-the-weaviate-function)
* [Related packages, products and repos](#related-packages-products-and-repos)
* [About different Weaviate versions](#about-different-weaviate-versions)
* [Contributing and Gitflow](#contributing-and-gitflow)
* [Main contributors](#main-contributors)
* [More info](#more-info)

### How does it work?
Google provides different libraries for interacting with the Weave protocol ([more info](http://weaviate.com/)). By changing the end-points to your own private cloud that runs Weaviate. You can use the complete Weave and Brillo software solutions within you own cloud solution.

### Installation
- Install Node version >5.0.0 ([more info](https://nodejs.org/en/download/package-manager))
- Install Weaviate from NPM: `npm install weaviate --save`
- Install a Cassandra database ([more info](https://www.digitalocean.com/community/tutorials/how-to-install-cassandra-and-run-a-single-node-cluster-on-ubuntu-14-04)) and import the CQL file found in the [Weaviate-Cassandra](https://github.com/weaviate/weaviate-cassandra) repo.
- In your server Javascript file, add the following:
```javascript
const weaviate = require('weaviate');
weaviate( CONFIGURATION OBJECT )
    .done((weaveObject) => { CALLBACK });
```

### Using the weaviate() function
The weaviate function needs configuration objects and returns an optional promise with a weaveObject.

**Configuration Object example**:
```javascript
{
	https		: false, // set to false = http usage, true = https usage
	https_opts	: {      // standard Express https options
        key: key,
        cert: cert
	},
	dbHostname  : 'localhost', // Cassandra hostname
	dbPort 		: 1000,        // Cassandra port
	dbName 		: 'test',      // Cassandra db name
	dbPassword 	: 'abc',       // Cassandra password
	dbContactpoints : ['h1'],      // Cassandra contain points
	hostname 	: 'localhost', // hostname for the service
	port 	 	: 8080,        // port for the service
	formatIn 	: 'JSON',      // JSON or CBOR (note: experimental)
	formatOut 	: 'JSON',      // JSON or CBOR (note: experimental)
	stdoutLog 	: true         // log all usages via stdout
}
```

**Promise weaveObject example**:

The weaveObject contains information about the request that you can use for custom processing purposes.
* params = the provided URL params
* body = the provided body
* response = the response generated by Weaviate
* requestHeaders = the original request headers

**Complete example**:
```javascript
weaviate({
	https		: false,
	dbHostname 	: 'localhost',
	dbPort 		: 1000,
	dbName 		: 'test',
	dbPassword 	: 'abc',
	hostname 	: 'localhost',
	port 	 	: '8080',
	formatIn 	: 'JSON',
	formatOut 	: 'JSON',
	stdoutLog 	: true
})
.done((weaveObject) => {
	/**
	 * The request is done
	 */
	console.log(weaveObject);
});
```

### Related packages, products and repos
There are a few related packages. More will be added soon.
- [Weaviate Cassandra](https://github.com/weaviate/weaviate-cassandra) contains the CQL file for Cassandra.
- [Weaviate Console](https://github.com/weaviate/weaviate-console) is a front-end console you might want to use to interact with the APIs.
- [Weaviate Auth](https://github.com/weaviate/weaviate-auth) is the authentication protocol.

### About different Weaviate versions
Weaviate comes in three versions.
* **Community version**: Open Source
* **Enterprise hosted version**: a multinode hosted version. [More info on weaviate.com](http://weaviate.com)
* **Enterprise version**: an SLA based version. [More info on weaviate.com](http://weaviate.com)

For more information, please contact: yourfriends@weaviate.com

### Contributing and Gitflow
Read more in the [CONTRIBUTE.md](CONTRIBUTE.md) file.

### About Weaviate
[Weave](https://developers.google.com/weave) is a high quality, open source end-to-end communications platform for IoT that allows you to connect and manage devices in a generic way. We understand that sometimes you need to be in control over your complete dataset. This may depend on the nature of your business, on analytics and predictions you want to do, or because you want to extend the protocol. Weaviate works as a replicate of the Google Weave cloud and runs on your own cloud or datacenter solution.

### Main contributors
- Bob van Luijt (bob@weaviate.com, @bobvanluijt)

### More info
Please keep checking back this repo or the website, we will start publishing software soon.

- Website: http://www.weaviate.com
- Mailing: http://eepurl.com/bRsMir
- Twitter: http://twitter.com/weaviate_iot
