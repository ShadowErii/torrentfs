# P2P file system of cortex full node
```
go get github.com/CortexFoundation/torrentfs

or

make
```
```cd build/bin```

```mkdir -p store/f3bc013c3cda4bb9f74a6f7696e66efa7be92a9a```

```echo "Hello torrent" > store/f3bc013c3cda4bb9f74a6f7696e66efa7be92a9a/data ```

#### torrent-create : to create torrent file
```./torrent-create store/f3bc013c3cda4bb9f74a6f7696e66efa7be92a9a/data > store/f3bc013c3cda4bb9f74a6f7696e66efa7be92a9a/torrent```
#### torrent-magnet : load info hash from torrent file
```./torrent-magnet < store/f3bc013c3cda4bb9f74a6f7696e66efa7be92a9a/torrent```

```
tree store

f3bc013c3cda4bb9f74a6f7696e66efa7be92a9a
├── data
└── torrent
```
#### seeding : to seed file to dht
```./seeding -dataDir=store```
#### torrent : to download file
```./torrent download 'infohash:f3bc013c3cda4bb9f74a6f7696e66efa7be92a9a' ```
