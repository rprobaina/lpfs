# lpfs
`lpfs` is a go library that makes it easier to get `procfs` info on a Linux system.


# How to use it
1. Get the `lpfs` module.

   ```bash
   $ go get github.com/rprobaina/lpfs
   ```

2. Import `lpfs` and use the functions.

   ```bash
   ...
   import (
     "github.com/rprobaina/lpfs"
   )
   
   your_var, err := lpfs.Function()
   
   uptime, err : GetUptimeSystem() // A real example
   ```

3. See the available functions, by using `go doc`.

   ```bash
   $ go doc -all lpfs
   ```

4. You can find some usage samples [here](https://github.com/rprobaina/lpfs/tree/main/examples).
