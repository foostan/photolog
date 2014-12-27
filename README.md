# Photolog
Tools for managing photos, include cli tools and web UI.

## Relocate command
Relocate photos by datetime and rename to the hash of photo information(hash params of `PhotoInfo struct`).

### Usage
```
NAME:
   relocate - Relocate photos

USAGE:
   command relocate [command options] [arguments...]

DESCRIPTION:
   Relocate photos by datetime and rename to the hash of photo information

OPTIONS:
   --basepath '.'	base directory path of target files
   --log-level 'warn'	logger level
```

## Contributing
1. Fork it ( https://github.com/[my-github-username]/photolog/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request