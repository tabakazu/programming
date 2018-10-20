# Chef Zero
## Initializes remote host
Install chef-client to node.
```bash
$ bundle exec knife zero bootstrap <ip> --node-name <nodename>
```

## Check node info
Check node details.
```bash
$ bundle exec knife node show <nodename>
```

Check node list.
```bash
$ bundle exec knife node list
```

## Run chef-client to node
Run chef-client to node. (why-run mode)
```bash
$ bundle exec knife zero converge "name:*" --why-run
```

Run chef-client to node.
```bash
$ bundle exec knife zero converge "name:*"
```

## Run chef-client to node (Single mode)
Run specified recipe for a single run.
```bash
$ bundle exec knife zero converge "name:<nodename>" -o recipe[base::packages]
```