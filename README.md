# goinventory

Cloud agnostic inventory system currently supporting:

* AWS

Aims to support:

* OVH
* GCP
* Scaleway
* Leaseweb

 This is aims to be a clone of [linkfluence/inventory](https://github.com/linkfluence/inventory) in golang

* Api aims to be compatible with linkfluence/inventory
* yaml as configuration file (use viper library) instead of clj for service configuration

## Storage

### Non Atomic

Inventory are basically stored in flat yaml files so it can be read easily, however for distributed setup the following storage are/will be supported:

* yaml stored into aws s3
* yaml stored into gcp storage

Maybe:

* oss (aliyun object store)
* swift

### Atomic

Since this a full rewrite atomic storage will be implemented and aims to support consul/postgres/mysql

### Replication

Storage replication will be implemented between atomic and non atomic. The design is still single writer/multiple reader.

## Usage

Building

    $ go build

Starting

    $ ./goinventory

## Documentation

Documentation is available [here](doc/intro.md)

## General stuffs

Copyright Jean-Baptiste Besselat © 2021 Adot SAS

Distributed under the Eclipse Public License either version 1.0 or (at your option) any later version.
