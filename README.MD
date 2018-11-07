# HANA Utility

## WHAT

This command line consist of HANA specific utilities that can be used by Administrator and HANA developer (esp. HANA modeller) to help them get some specific information (which usually can be queried using SQL)

The tools consist of:
1. Admin
   - List Active objects
   - List Inactive objects
   - Get Memory Consumption
   - List Deactivated users
   - List User login information
2. Dev
   - View dependencies what Tables used by certain Packages.
   - View dependencies what Models are using specific Tables.
   - List out Models within certain Package and what Tables are used within those Models.

## WHY

I have a lot of queries regarding HANA information due to my unspecific job of BI. I tired to query again and again which someone usually can simply run certain SQL on sql console to get those information.

## USAGE

1. Copy/Rename config.ini.sample become config.ini
2. Edit config.ini, fill in the respecting information inside this file and save
```ini
host = 10.0.0.1
uid = SAMPLEUSER
pwd = Y0urP@ssw0rd
port = 30015
```
3. Compile this project
4. Run the binary

## THANKS TO

- github.com/SAP/go-hdb/driver (special thanks to SAP to bring HANA driver to Go)
- github.com/spf13/cobra
- github.com/go-ini/ini