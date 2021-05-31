# Readmine Time Checker

Developers' daily report for managers.

telegram bot written on Go

## Environment Variables

USERS_LIST - serialized array of managers telegram channels and developers

BOT_TOKEN - telegram bot token

REDMINE_URL - link to redmine

REDMINE_API_KEY - redmine API key

## example of USER_LIST env in json

```json
[
  {
    "manager_channel_id": __manager_chanel_id__,
    "developers_ids": [
      1,
      2,
      3
    ]
  }
]
```
*you should serialize it before you put it in the .env file

## License

Copyright (c) 2021

Licensed under the [MIT license](LICENSE).
