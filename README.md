This is just a basic tool to convert the output from Microsofts 'csvde' into a format that can be imported into MySQL.

Namely, it's job is to detect the objectGUID field and strip the screwed up escaping insists on using.

But more than anything else, this is just a good excuse to learn some more Go skills.

Was thinking of going all the way and creating a resulting SQL create query but it's going to be quicker to use PhpMyAdmin's tools for this particular job.
