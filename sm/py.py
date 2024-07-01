#!/usr/bin/python
# -*- coding: utf-8 -*-

import sys


class Lookup(object):
    ttl = 30

    def __init__(self, query):
        (_type, qname, qclass, qtype, _id, ip) = query

        self.has_result = False
        qname_lower = qname.lower()

        self.results = []

        self.results.append('DATA\t%s\t%s\t%s\t%d\t-1\t%s' % (
            qname_lower, qclass, 'A', 300, '192.168.40.223'))
        self.has_result = True

    def str_result(self):
        if self.has_result:
            return '\n'.join(self.results)
        else:
            return ''


class DNSbackend(object):

    def __init__(self, filein, fileout):
        self.filein = filein
        self.fileout = fileout

        self._process_requests()

    def _fprint(self, message):
        self.fileout.write(message + '\n')
        self.fileout.flush()

    def _process_requests(self):
        first_time = True

        while 1:
            rawline = self.filein.readline()

            if rawline == '':
                return

            line = rawline.rstrip()

            if first_time:
                if line == 'HELO\t1':
                    self._fprint('OK\tPython backend ready.')
                else:
                    rawline = self.filein.readline()
                    sys.exit(1)
                first_time = False
            else:
                # print("-------------", line)
                query = line.split('\t')
                if len(query) != 6:
                    self._fprint('LOG\tPowerDNS sent un-parseable line')
                    self._fprint('FAIL')
                else:
                    lookup = Lookup(query)
                    if lookup.has_result:
                        pdns_result = lookup.str_result()
                        self._fprint(pdns_result)
                    self._fprint('END')


if __name__ == "__main__":
    infile = sys.stdin
    outfile = sys.stdout

    try:
        DNSbackend(infile, outfile)

    except:
        raise
