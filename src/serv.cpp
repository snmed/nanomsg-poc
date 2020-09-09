
#include <string> // use for windows getline
#include <iostream>
#include <nng/nng.h>
#include <nng/protocol/reqrep0/rep.h>
#include <nng/protocol/reqrep0/req.h>

using namespace std;

void fatal(const char *func, int rv)
{
    fprintf(stderr, "%s: %s\n", func, nng_strerror(rv));
    exit(1);
}

int main()
{
    cout << "Welcome to NNG example cli..." << endl;
    const char *url = "ipc:///tmp/test.ipc";
    nng_socket sock;
    int rv;

    if ((rv = nng_req0_open(&sock)) != 0)
    {
        fatal("nng_req0_open", rv);
    }

    if ((rv = nng_dial(sock, url, NULL, 0)) != 0)
    {
        fatal("nng_dial", rv);
    }
    string msg = "";
    size_t buf_sz = 1024 * 1024 * 2; // 2Mb
    char *buf = new char[buf_sz];

    for (;;)
    {
        cout << "Enter your name: ";
        getline(cin, msg);

        // Send msg to server
        auto data = msg.c_str();

        rv = nng_send(sock, (void *)data, strlen(data), 0);
        if (rv != 0)
        {
            cout << "failed to send data: " << nng_strerror(rv) << endl;
            continue;
        }

        size_t sz = buf_sz;
        memset(buf, 0, sz);
        rv = nng_recv(sock, buf, &sz, 0);
        if (rv != 0)
        {
            cout << "failed to receive data: " << nng_strerror(rv) << endl;
            continue;
        }

        cout << buf << endl;
    }
    delete[] buf;
}