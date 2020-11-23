package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/theme"
)

// Icons Has the list of icons from the standard theme.
// ToDo: Will have to look for a way to sync the list from `fyne_demo`
var icons = map[string]fyne.Resource{
	"CancelIcon":        theme.CancelIcon(),
	"ConfirmIcon":       theme.ConfirmIcon(),
	"DeleteIcon":        theme.DeleteIcon(),
	"SearchIcon":        theme.SearchIcon(),
	"SearchReplaceIcon": theme.SearchReplaceIcon(),

	"CheckButtonIcon":        theme.CheckButtonIcon(),
	"CheckButtonCheckedIcon": theme.CheckButtonCheckedIcon(),
	"RadioButtonIcon":        theme.RadioButtonIcon(),
	"RadioButtonCheckedIcon": theme.RadioButtonCheckedIcon(),

	"ColorAchromaticIcon": theme.ColorAchromaticIcon(),
	"ColorChromaticIcon":  theme.ColorChromaticIcon(),
	"ColorPaletteIcon":    theme.ColorPaletteIcon(),

	"ContentAddIcon":    theme.ContentAddIcon(),
	"ContentRemoveIcon": theme.ContentRemoveIcon(),
	"ContentClearIcon":  theme.ContentClearIcon(),
	"ContentCutIcon":    theme.ContentCutIcon(),
	"ContentCopyIcon":   theme.ContentCopyIcon(),
	"ContentPasteIcon":  theme.ContentPasteIcon(),
	"ContentRedoIcon":   theme.ContentRedoIcon(),
	"ContentUndoIcon":   theme.ContentUndoIcon(),

	"InfoIcon":     theme.InfoIcon(),
	"ErrorIcon":    theme.ErrorIcon(),
	"QuestionIcon": theme.QuestionIcon(),
	"WarningIcon":  theme.WarningIcon(),

	"DocumentIcon":       theme.DocumentIcon(),
	"DocumentCreateIcon": theme.DocumentCreateIcon(),
	"DocumentPrintIcon":  theme.DocumentPrintIcon(),
	"DocumentSaveIcon":   theme.DocumentSaveIcon(),

	"FileIcon":            theme.FileIcon(),
	"FileApplicationIcon": theme.FileApplicationIcon(),
	"FileAudioIcon":       theme.FileAudioIcon(),
	"FileImageIcon":       theme.FileImageIcon(),
	"FileTextIcon":        theme.FileTextIcon(),
	"FileVideoIcon":       theme.FileVideoIcon(),
	"FolderIcon":          theme.FolderIcon(),
	"FolderNewIcon":       theme.FolderNewIcon(),
	"FolderOpenIcon":      theme.FolderOpenIcon(),
	"ComputerIcon":        theme.ComputerIcon(),
	"HomeIcon":            theme.HomeIcon(),
	"HelpIcon":            theme.HelpIcon(),
	"HistoryIcon":         theme.HistoryIcon(),
	"SettingsIcon":        theme.SettingsIcon(),
	"StorageIcon":         theme.StorageIcon(),
	"DownloadIcon":        theme.DownloadIcon(),
	// "UploadIcon":          theme.UploadIcon(),

	"ViewFullScreenIcon": theme.ViewFullScreenIcon(),
	"ViewRestoreIcon":    theme.ViewRestoreIcon(),
	"ViewRefreshIcon":    theme.ViewRefreshIcon(),
	"VisibilityIcon":     theme.VisibilityIcon(),
	"VisibilityOffIcon":  theme.VisibilityOffIcon(),
	"ZoomFitIcon":        theme.ZoomFitIcon(),
	"ZoomInIcon":         theme.ZoomInIcon(),
	"ZoomOutIcon":        theme.ZoomOutIcon(),

	"MoveDownIcon": theme.MoveDownIcon(),
	"MoveUpIcon":   theme.MoveUpIcon(),

	"NavigateBackIcon": theme.NavigateBackIcon(),
	"NavigateNextIcon": theme.NavigateNextIcon(),

	"Menu":         theme.MenuIcon(),
	"MenuExpand":   theme.MenuExpandIcon(),
	"MenuDropDown": theme.MenuDropDownIcon(),
	"MenuDropUp":   theme.MenuDropUpIcon(),

	"MailAttachmentIcon": theme.MailAttachmentIcon(),
	"MailComposeIcon":    theme.MailComposeIcon(),
	"MailForwardIcon":    theme.MailForwardIcon(),
	"MailReplyIcon":      theme.MailReplyIcon(),
	"MailReplyAllIcon":   theme.MailReplyAllIcon(),
	"MailSendIcon":       theme.MailSendIcon(),

	"MediaFastForward": theme.MediaFastForwardIcon(),
	"MediaFastRewind":  theme.MediaFastRewindIcon(),
	"MediaPause":       theme.MediaPauseIcon(),
	"MediaPlay":        theme.MediaPlayIcon(),
	// "MediaStop":         theme.MediaStopIcon(),
	"MediaRecord":       theme.MediaRecordIcon(),
	"MediaReplay":       theme.MediaReplayIcon(),
	"MediaSkipNext":     theme.MediaSkipNextIcon(),
	"MediaSkipPrevious": theme.MediaSkipPreviousIcon(),

	"VolumeDown": theme.VolumeDownIcon(),
	"VolumeMute": theme.VolumeMuteIcon(),
	"VolumeUp":   theme.VolumeUpIcon(),
}

// IconNames Gets the list of the names of the icons available
var iconNames = make([]string, len(icons))

// GenerateRepo creates all the repo variables. For example, Adds all the icon names to the IconNames variable
func generateRepo() {
	i := 0
	for k := range icons {
		iconNames[i] = k
		i++
	}
}
