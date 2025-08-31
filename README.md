# GoConfiguration
Configuration Manager for Go projects

This package makes loading configurations as easy as it can be.
---
GoConfiguration package supports: 
- Reading configurations from files 
   - [x] json
   - [ ] xml
   - [ ] yaml
   - [ ] toml
   - [ ] env
- [ ] Reading configurations from environment variables
- [ ] Reading configurations from flags
- [ ] configuration pipeline
   - [ ] using multiple configuration sources
   - [ ] configuration override
   - [ ] optional configuration files
- configuration hot reloading 
   - [ ] hot reloading configuration files
   - [ ] hot reloading environment variables

- [x] getting configuration sections (as dictionary)
   - [ ] fast section selection (using `.` as key)
- [x] getting configuration as an struct
