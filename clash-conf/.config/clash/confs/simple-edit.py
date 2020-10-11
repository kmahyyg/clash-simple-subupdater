#!/usr/bin/python3
# -*- encoding: utf-8 -*-

import yaml
import sys

# Read config file in

try:
    ispname = sys.argv[1].lower()
    ftname = sys.argv[2].lower()
except IndexError:
    print("Append ISP Name to Serve. And Config Name to Save.")
    sys.exit(1)

fd = open(ispname+'.yaml', 'r')
servConf = yaml.safe_load(fd)

# Modify the port

servConf['mixed-port'] = 17893
servConf['redir-port'] = 17894

try:
    del servConf['port']
    del servConf['socks-port']
except:
    pass

# Allow Lan Connection

servConf['allow-lan'] = True
servConf['mode'] = 'rule'
servConf['ipv6'] = False
servConf['bind-address'] = '*'

# Allow Remote Conf
servConf['external-controller'] = '0.0.0.0:17900'
servConf['external-ui'] = 'webui'

# Set Global Secret
servConf['secret'] = 'A8a9B0b6c0c4D'

# Allow DNS
servConf['dns']['enable'] = True
servConf['dns']['ipv6'] = False
servConf['dns']['enhanced-mode'] = 'fake-ip'
servConf['dns']['listen'] = '0.0.0.0:10053'
servConf['dns']['fake-ip-range'] = '198.18.0.1/16'
servConf['dns']['use-hosts'] = True

# Ban BT
servConf['rules'].insert(1,'DST-PORT,6881,REJECT')

# Save out
finaldt = yaml.dump(servConf, indent=2, allow_unicode=True, encoding='utf-8', default_flow_style=False, sort_keys=False)
with open(ftname+'.yaml', 'wb') as ffd:
    ffd.write(finaldt)

