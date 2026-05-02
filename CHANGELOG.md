# Changelog
All notable changes to Inkstamp will be documented here

## [0.6.0] - 2026-05-02

### Added
- Made `https://github.com/phamio/termcolor` internal in inkstamp

### Changed
- Supported golang version from `go 1.25.2` to `go 1.25.0`


## [0.5.1] - 2026-04-28 

### Added
- Center alignment ('^') for placeholders
- Custom fill characters (`[0:<20:.]`, `[0:>10:=]`)
- `FillChar`, `Width`, `Align` fields in `TempPart` (replaces `FormatStr`)
- Placeholder now supports fill characters with new syntax as `[index:alignmentWidth:fillCharacter]`

### Changed
- Renamed from Crayon to Inkstamp
- Repository moved to `https://github.com/inkstamp/inkstamp`

### Removed
- `FormatStr` field from internal `TempPart` struct


## [0.5.0] - 2026-04-24

### Added
- True color to 256 color palette fallback
- Full color fallback chain (truecolor ==> 256 ==> 16)
- Changed syntax of light colors (lightred ==> lred)
- Dumb terminal detection

### Changed
- Hex validation now properly requires '#' prefix
- Fixed parse256ColorCode undefined variable bug
- Improved parseRGB length and int validation

### Fixed
- Unclosed bracket handling in templates


## [0.4.0] - 2026-04-06

### Added
- Inline padding in placeholders

---

## [0.3.0] - 2026-04-06

### Added
- Color support on windows
