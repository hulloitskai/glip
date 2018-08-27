package glip

import "github.com/steven-xie/glip/portal"

// PShellBoard is capable of interacting with the Windows clipboard using
// the "Get-Clipboard" and "Set-Clipboard" PowerShell cmdlets.
//
// See the documentation for the aforementioned cmdlets at
// https://bit.ly/2wqjBzI.
type PShellBoard struct {
	// Format represents Get-Clipboard's "-Format" flag setting.
	Format PSBFormat

	// TextFormatType represents Get-Clipboard's "-TextFormatType" flag setting.
	TextFormatType PSBTextFormatType

	// Raw represents Get-Clipboard's "-Raw" flag setting.
	//
	// If enabled, Get-Clipboard will ignore newline characters and get the
	// entire contents of the clipboard.
	Raw bool

	// Append represents Set-Clipboard's "-Append" flag setting.
	//
	// If enabled, Set-Clipboard will append data to the clipboard rather than
	// overwrite its previous contents.
	Append bool

	// AsHTML represents Set-Clipboard's "-AsHtml" flag setting.
	//
	// If enabled, Set-Clipboard will render content sent to the clipboard as
	// HTML.
	AsHTML bool

	*dualBoard
}

// PSBFormat is a string that represents a PShellBoard "-Format" flag option.
type PSBFormat string

// PShellBoard "-Format" flag options.
const (
	PSBText         PSBFormat = "Text"
	PSBFileDropList           = "FileDropList"
	PSBImage                  = "Image"
	PSBAudio                  = "Audio"
)

// PSBTextFormatType represents a PShellBoard "-TextFormatType" flag option.
type PSBTextFormatType string

// PShellBoard "-TextFormatType" flag options.
const (
	PSBTextType        PSBTextFormatType = "Text"
	PSBUnicodeTextType                   = "UnicodeText"
	PSBRtfType                           = "Rtf"
	PSBHtmlType                          = "Html"
	PSBCsvType                           = "CommaSeparatedValue"
)

// cmdletPortal makes a portal.Portal from a PowerShell cmdlet.
func cmdletPortal(name string) *portal.Portal {
	return portal.New("PowerShell", "-Command", name)
}

// NewPShellBoard creates a new PShellBoard with zeroed (default) flag options,
// if its underlying programs can be found in the system path.
func NewPShellBoard() (psb *PShellBoard, err error) {
	if err = ensureCmdExists("PowerShell"); err != nil {
		return nil, err
	}

	ps := &PShellBoard{
		dualBoard: &dualBoard{
			Writer: newDynPortal("Set-Clipboard"),
			Reader: newDynPortal("Get-Clipboard"),
		},
	}

	ps.Writer.GetArgs = ps.generateWriterArgs
	ps.Reader.GetArgs = ps.generateReaderArgs
	return ps, nil
}

func (psb *PShellBoard) generateReaderArgs() []string {
	args := make([]string, 0, 3)

	if psb.Raw {
		args = append(args, "-Raw")
	}
	if psb.Format != "" {
		args = append(args, "-Format", string(psb.Format))
	}
	if psb.TextFormatType != "" {
		args = append(args, "-TextFormatType", string(psb.TextFormatType))
	}

	return args
}

func (psb *PShellBoard) generateWriterArgs() []string {
	args := make([]string, 0, 2)

	if psb.Append {
		args = append(args, "-Append")
	}
	if psb.AsHTML {
		args = append(args, "-AsHtml")
	}

	return args
}
