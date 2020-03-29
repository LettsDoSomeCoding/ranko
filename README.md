# Ranko

## Docker Test Environment Setup

Ranko is a containerised application using Docker. As such a test environment needs to be set up in order to test that Ranko correctly works inside of a container. Here are the instructions to set up that test environment:

1. Install VirtualBox v6.0.14+ ([Download from here](https://www.virtualbox.org/wiki/Downloads))
2. Install Vagrant v2.2.7+ ([Download from here](https://www.vagrantup.com/downloads.html))
3. Inside the Ranko code directory run `vagrant up`
4. Once the above command has finished outputting to your command prompt, run `vagrant ssh` to connect into your test environment

```NOTE: On first install of this environment it will take a long time to download the necessary VM image but this won't be required for subsequent installations```

The code for Ranko is stored in the `/vagrant` folder in the test environment. When you are finished doing your docker testing disconnect from the environment with `Ctrl + d` and then remove the environment with `vagrant destroy`.

To run Ranko:
1.  In the root module directory (the one containing go.mod), run `go build`. This will create an executable called `ranko` (or `ranko.exe` on Windows)
2.  Run the executable
