// +build darwin

package glip

// DarwinBoard is capable of interacting with the macOS clipboard ("pasteboard")
// using "pbcopy" and "pbpaste".
type DarwinBoard struct {
	// PBoard represents pbcopy and pbpaste's "-pboard" flag setting. It defaults
	// to DPBGeneral.
	//
	// DarwinBoard will read / write data into this DPBoard.
	PBoard DPBoard

	// Prefer represents pbpaste's "-Prefer" flag setting.
	//
	// This value corresponds with the type of data that pbpaste will look
	// for in the clipboard.
	Prefer DPrefer

	*dualBoard
}

// DPBoard is a string that identifies a particular macOS ("darwin") pasteboard.
type DPBoard string

// These DPBoards corresponding to macOS ("darwin") pasteboards.
const (
	DPBGeneral DPBoard = "general" // the default macOS pasteboard
	DPBRuler           = "ruler"
	DPBFind            = "find"
	DPBFont            = "font"
)

// DPrefer represents the value of the "-Prefer" flag setting for pbpaste.
type DPrefer string

// These are possible settings of the "-Prefer" flag for pbpaste.
const (
	PreferTxt DPrefer = "txt"
	PreferRTF         = "rtf"
	PreferPS          = "ps"
)

// NewDarwinBoard creates a new default DarwinBoard instance, if its underlying
// programs can be found in the system path.
//
// It targets the "general" macOS pasteboard.
func NewDarwinBoard() (db *DarwinBoard, err error) {
	return NewDarwinBoardTarget(DPBGeneral)
}

// NewDarwinBoardTarget creates a new DarwinBoard targetting a specific macOS
// board, if its underlying programs can be found in the system path.
func NewDarwinBoardTarget(board DPBoard) (db *DarwinBoard, err error) {
	const (
		copycmd  = "pbcopy"
		pastecmd = "pbpaste"
	)

	if err = ensureCmdExists(copycmd); err != nil {
		return nil, err
	}
	if err = ensureCmdExists(pastecmd); err != nil {
		return nil, err
	}

	pb := &DarwinBoard{
		dualBoard: newDualBoard(
			newDynPortal(copycmd),
			newDynPortal(pastecmd),
		),
		PBoard: board,
	}

	pb.Writer.GetArgs = pb.generateWriterArgs
	pb.Reader.GetArgs = pb.generateReaderArgs
	return pb, nil
}

func (pb *DarwinBoard) generateWriterArgs() []string {
	return []string{"-pboard", string(pb.PBoard)}
}

func (pb *DarwinBoard) generateReaderArgs() []string {
	args := pb.generateWriterArgs()
	if pb.Prefer != "" {
		args = append(args, "-Prefer", string(pb.Prefer))
	}
	return args
}
