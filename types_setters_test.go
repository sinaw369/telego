package telego

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplyKeyboardMarkup_Setters(t *testing.T) {
	r := (&ReplyKeyboardMarkup{}).
		WithKeyboard([][]KeyboardButton{{}}...).
		WithResizeKeyboard().
		WithOneTimeKeyboard().
		WithInputFieldPlaceholder("InputFieldPlaceholder").
		WithSelective()

	assert.Equal(t, &ReplyKeyboardMarkup{
		Keyboard:              [][]KeyboardButton{{}},
		ResizeKeyboard:        true,
		OneTimeKeyboard:       true,
		InputFieldPlaceholder: "InputFieldPlaceholder",
		Selective:             true,
	}, r)
}

func TestKeyboardButton_Setters(t *testing.T) {
	k := (KeyboardButton{}).
		WithText("Text").
		WithRequestContact().
		WithRequestLocation().
		WithRequestPoll(&KeyboardButtonPollType{Type: "RequestPoll"}).
		WithWebApp(&WebAppInfo{})

	assert.Equal(t, KeyboardButton{
		Text:            "Text",
		RequestContact:  true,
		RequestLocation: true,
		RequestPoll:     &KeyboardButtonPollType{Type: "RequestPoll"},
		WebApp:          &WebAppInfo{},
	}, k)
}

func TestReplyKeyboardRemove_Setters(t *testing.T) {
	r := (&ReplyKeyboardRemove{}).
		WithRemoveKeyboard().
		WithSelective()

	assert.Equal(t, &ReplyKeyboardRemove{
		RemoveKeyboard: true,
		Selective:      true,
	}, r)
}

func TestInlineKeyboardMarkup_Setters(t *testing.T) {
	i := (&InlineKeyboardMarkup{}).
		WithInlineKeyboard([][]InlineKeyboardButton{{}}...)

	assert.Equal(t, &InlineKeyboardMarkup{
		InlineKeyboard: [][]InlineKeyboardButton{{}},
	}, i)
}

func TestInlineKeyboardButton_Setters(t *testing.T) {
	i := (InlineKeyboardButton{}).
		WithText("Text").
		WithURL("URL").
		WithCallbackData("CallbackData").
		WithWebApp(&WebAppInfo{}).
		WithLoginURL(&LoginURL{URL: "LoginURL"}).
		WithSwitchInlineQuery("SwitchInlineQuery").
		WithSwitchInlineQueryCurrentChat("SwitchInlineQueryCurrentChat").
		WithCallbackGame(&CallbackGame{}).
		WithPay()

	assert.Equal(t, InlineKeyboardButton{
		Text:                         "Text",
		URL:                          "URL",
		CallbackData:                 "CallbackData",
		WebApp:                       &WebAppInfo{},
		LoginURL:                     &LoginURL{URL: "LoginURL"},
		SwitchInlineQuery:            "SwitchInlineQuery",
		SwitchInlineQueryCurrentChat: "SwitchInlineQueryCurrentChat",
		CallbackGame:                 &CallbackGame{},
		Pay:                          true,
	}, i)
}

func TestForceReply_Setters(t *testing.T) {
	f := (&ForceReply{}).
		WithForceReply().
		WithInputFieldPlaceholder("InputFieldPlaceholder").
		WithSelective()

	assert.Equal(t, &ForceReply{
		ForceReply:            true,
		InputFieldPlaceholder: "InputFieldPlaceholder",
		Selective:             true,
	}, f)
}

func TestMenuButtonWebApp_Setters(t *testing.T) {
	m := (&MenuButtonWebApp{}).
		WithText("Text").
		WithWebApp(WebAppInfo{})

	assert.Equal(t, &MenuButtonWebApp{
		Text:   "Text",
		WebApp: WebAppInfo{},
	}, m)
}

func TestInputMediaPhoto_Setters(t *testing.T) {
	i := (&InputMediaPhoto{}).
		WithMedia(testInputFile).
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...)

	assert.Equal(t, &InputMediaPhoto{
		Media:           testInputFile,
		Caption:         "Caption",
		ParseMode:       "ParseMode",
		CaptionEntities: []MessageEntity{{Type: "CaptionEntities"}},
	}, i)
}

func TestInputMediaVideo_Setters(t *testing.T) {
	i := (&InputMediaVideo{}).
		WithMedia(testInputFile).
		WithThumb(&testInputFile).
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithWidth(1).
		WithHeight(2).
		WithDuration(3).
		WithSupportsStreaming()

	assert.Equal(t, &InputMediaVideo{
		Media:             testInputFile,
		Thumb:             &testInputFile,
		Caption:           "Caption",
		ParseMode:         "ParseMode",
		CaptionEntities:   []MessageEntity{{Type: "CaptionEntities"}},
		Width:             1,
		Height:            2,
		Duration:          3,
		SupportsStreaming: true,
	}, i)
}

func TestInputMediaAnimation_Setters(t *testing.T) {
	i := (&InputMediaAnimation{}).
		WithMedia(testInputFile).
		WithThumb(&testInputFile).
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithWidth(1).
		WithHeight(2).
		WithDuration(3)

	assert.Equal(t, &InputMediaAnimation{
		Media:           testInputFile,
		Thumb:           &testInputFile,
		Caption:         "Caption",
		ParseMode:       "ParseMode",
		CaptionEntities: []MessageEntity{{Type: "CaptionEntities"}},
		Width:           1,
		Height:          2,
		Duration:        3,
	}, i)
}

func TestInputMediaAudio_Setters(t *testing.T) {
	i := (&InputMediaAudio{}).
		WithMedia(testInputFile).
		WithThumb(&testInputFile).
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithDuration(1).
		WithPerformer("Performer").
		WithTitle("Title")

	assert.Equal(t, &InputMediaAudio{
		Media:           testInputFile,
		Thumb:           &testInputFile,
		Caption:         "Caption",
		ParseMode:       "ParseMode",
		CaptionEntities: []MessageEntity{{Type: "CaptionEntities"}},
		Duration:        1,
		Performer:       "Performer",
		Title:           "Title",
	}, i)
}

func TestInputMediaDocument_Setters(t *testing.T) {
	i := (&InputMediaDocument{}).
		WithMedia(testInputFile).
		WithThumb(&testInputFile).
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithDisableContentTypeDetection()

	assert.Equal(t, &InputMediaDocument{
		Media:                       testInputFile,
		Thumb:                       &testInputFile,
		Caption:                     "Caption",
		ParseMode:                   "ParseMode",
		CaptionEntities:             []MessageEntity{{Type: "CaptionEntities"}},
		DisableContentTypeDetection: true,
	}, i)
}

func TestInlineQueryResultArticle_Setters(t *testing.T) {
	i := (&InlineQueryResultArticle{}).
		WithID("ID").
		WithTitle("Title").
		WithInputMessageContent(&InputTextMessageContent{}).
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}}).
		WithURL("URL").
		WithHideURL().
		WithDescription("Description").
		WithThumbURL("ThumbURL").
		WithThumbWidth(1).
		WithThumbHeight(2)

	assert.Equal(t, &InlineQueryResultArticle{
		ID:                  "ID",
		Title:               "Title",
		InputMessageContent: &InputTextMessageContent{},
		ReplyMarkup:         &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
		URL:                 "URL",
		HideURL:             true,
		Description:         "Description",
		ThumbURL:            "ThumbURL",
		ThumbWidth:          1,
		ThumbHeight:         2,
	}, i)
}

func TestInlineQueryResultPhoto_Setters(t *testing.T) {
	i := (&InlineQueryResultPhoto{}).
		WithID("ID").
		WithPhotoURL("PhotoURL").
		WithThumbURL("ThumbURL").
		WithPhotoWidth(1).
		WithPhotoHeight(2).
		WithTitle("Title").
		WithDescription("Description").
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}}).
		WithInputMessageContent(&InputTextMessageContent{})

	assert.Equal(t, &InlineQueryResultPhoto{
		ID:                  "ID",
		PhotoURL:            "PhotoURL",
		ThumbURL:            "ThumbURL",
		PhotoWidth:          1,
		PhotoHeight:         2,
		Title:               "Title",
		Description:         "Description",
		Caption:             "Caption",
		ParseMode:           "ParseMode",
		CaptionEntities:     []MessageEntity{{Type: "CaptionEntities"}},
		ReplyMarkup:         &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
		InputMessageContent: &InputTextMessageContent{},
	}, i)
}

//nolint:dupl
func TestInlineQueryResultGif_Setters(t *testing.T) {
	i := (&InlineQueryResultGif{}).
		WithID("ID").
		WithGifURL("GifURL").
		WithGifWidth(1).
		WithGifHeight(2).
		WithGifDuration(3).
		WithThumbURL("ThumbURL").
		WithThumbMimeType("ThumbMimeType").
		WithTitle("Title").
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}}).
		WithInputMessageContent(&InputTextMessageContent{})

	assert.Equal(t, &InlineQueryResultGif{
		ID:                  "ID",
		GifURL:              "GifURL",
		GifWidth:            1,
		GifHeight:           2,
		GifDuration:         3,
		ThumbURL:            "ThumbURL",
		ThumbMimeType:       "ThumbMimeType",
		Title:               "Title",
		Caption:             "Caption",
		ParseMode:           "ParseMode",
		CaptionEntities:     []MessageEntity{{Type: "CaptionEntities"}},
		ReplyMarkup:         &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
		InputMessageContent: &InputTextMessageContent{},
	}, i)
}

//nolint:dupl
func TestInlineQueryResultMpeg4Gif_Setters(t *testing.T) {
	i := (&InlineQueryResultMpeg4Gif{}).
		WithID("ID").
		WithMpeg4URL("Mpeg4URL").
		WithMpeg4Width(1).
		WithMpeg4Height(2).
		WithMpeg4Duration(3).
		WithThumbURL("ThumbURL").
		WithThumbMimeType("ThumbMimeType").
		WithTitle("Title").
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}}).
		WithInputMessageContent(&InputTextMessageContent{})

	assert.Equal(t, &InlineQueryResultMpeg4Gif{
		ID:                  "ID",
		Mpeg4URL:            "Mpeg4URL",
		Mpeg4Width:          1,
		Mpeg4Height:         2,
		Mpeg4Duration:       3,
		ThumbURL:            "ThumbURL",
		ThumbMimeType:       "ThumbMimeType",
		Title:               "Title",
		Caption:             "Caption",
		ParseMode:           "ParseMode",
		CaptionEntities:     []MessageEntity{{Type: "CaptionEntities"}},
		ReplyMarkup:         &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
		InputMessageContent: &InputTextMessageContent{},
	}, i)
}

func TestInlineQueryResultVideo_Setters(t *testing.T) {
	i := (&InlineQueryResultVideo{}).
		WithID("ID").
		WithVideoURL("VideoURL").
		WithMimeType("MimeType").
		WithThumbURL("ThumbURL").
		WithTitle("Title").
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithVideoWidth(1).
		WithVideoHeight(2).
		WithVideoDuration(3).
		WithDescription("Description").
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}}).
		WithInputMessageContent(&InputTextMessageContent{})

	assert.Equal(t, &InlineQueryResultVideo{
		ID:                  "ID",
		VideoURL:            "VideoURL",
		MimeType:            "MimeType",
		ThumbURL:            "ThumbURL",
		Title:               "Title",
		Caption:             "Caption",
		ParseMode:           "ParseMode",
		CaptionEntities:     []MessageEntity{{Type: "CaptionEntities"}},
		VideoWidth:          1,
		VideoHeight:         2,
		VideoDuration:       3,
		Description:         "Description",
		ReplyMarkup:         &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
		InputMessageContent: &InputTextMessageContent{},
	}, i)
}

func TestInlineQueryResultAudio_Setters(t *testing.T) {
	i := (&InlineQueryResultAudio{}).
		WithID("ID").
		WithAudioURL("AudioURL").
		WithTitle("Title").
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithPerformer("Performer").
		WithAudioDuration(1).
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}}).
		WithInputMessageContent(&InputTextMessageContent{})

	assert.Equal(t, &InlineQueryResultAudio{
		ID:                  "ID",
		AudioURL:            "AudioURL",
		Title:               "Title",
		Caption:             "Caption",
		ParseMode:           "ParseMode",
		CaptionEntities:     []MessageEntity{{Type: "CaptionEntities"}},
		Performer:           "Performer",
		AudioDuration:       1,
		ReplyMarkup:         &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
		InputMessageContent: &InputTextMessageContent{},
	}, i)
}

func TestInlineQueryResultVoice_Setters(t *testing.T) {
	i := (&InlineQueryResultVoice{}).
		WithID("ID").
		WithVoiceURL("VoiceURL").
		WithTitle("Title").
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithVoiceDuration(1).
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}}).
		WithInputMessageContent(&InputTextMessageContent{})

	assert.Equal(t, &InlineQueryResultVoice{
		ID:                  "ID",
		VoiceURL:            "VoiceURL",
		Title:               "Title",
		Caption:             "Caption",
		ParseMode:           "ParseMode",
		CaptionEntities:     []MessageEntity{{Type: "CaptionEntities"}},
		VoiceDuration:       1,
		ReplyMarkup:         &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
		InputMessageContent: &InputTextMessageContent{},
	}, i)
}

func TestInlineQueryResultDocument_Setters(t *testing.T) {
	i := (&InlineQueryResultDocument{}).
		WithID("ID").
		WithTitle("Title").
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithDocumentURL("DocumentURL").
		WithMimeType("MimeType").
		WithDescription("Description").
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}}).
		WithInputMessageContent(&InputTextMessageContent{}).
		WithThumbURL("ThumbURL").
		WithThumbWidth(1).
		WithThumbHeight(2)

	assert.Equal(t, &InlineQueryResultDocument{
		ID:                  "ID",
		Title:               "Title",
		Caption:             "Caption",
		ParseMode:           "ParseMode",
		CaptionEntities:     []MessageEntity{{Type: "CaptionEntities"}},
		DocumentURL:         "DocumentURL",
		MimeType:            "MimeType",
		Description:         "Description",
		ReplyMarkup:         &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
		InputMessageContent: &InputTextMessageContent{},
		ThumbURL:            "ThumbURL",
		ThumbWidth:          1,
		ThumbHeight:         2,
	}, i)
}

func TestInlineQueryResultLocation_Setters(t *testing.T) {
	i := (&InlineQueryResultLocation{}).
		WithID("ID").
		WithTitle("Title").
		WithLivePeriod(1).
		WithHeading(2).
		WithProximityAlertRadius(3).
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}}).
		WithInputMessageContent(&InputTextMessageContent{}).
		WithThumbURL("ThumbURL").
		WithThumbWidth(4).
		WithThumbHeight(5)

	assert.Equal(t, &InlineQueryResultLocation{
		ID:                   "ID",
		Title:                "Title",
		LivePeriod:           1,
		Heading:              2,
		ProximityAlertRadius: 3,
		ReplyMarkup:          &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
		InputMessageContent:  &InputTextMessageContent{},
		ThumbURL:             "ThumbURL",
		ThumbWidth:           4,
		ThumbHeight:          5,
	}, i)
}

func TestInlineQueryResultVenue_Setters(t *testing.T) {
	i := (&InlineQueryResultVenue{}).
		WithID("ID").
		WithTitle("Title").
		WithAddress("Address").
		WithFoursquareID("FoursquareID").
		WithFoursquareType("FoursquareType").
		WithGooglePlaceID("GooglePlaceID").
		WithGooglePlaceType("GooglePlaceType").
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}}).
		WithInputMessageContent(&InputTextMessageContent{}).
		WithThumbURL("ThumbURL").
		WithThumbWidth(1).
		WithThumbHeight(2)

	assert.Equal(t, &InlineQueryResultVenue{
		ID:                  "ID",
		Title:               "Title",
		Address:             "Address",
		FoursquareID:        "FoursquareID",
		FoursquareType:      "FoursquareType",
		GooglePlaceID:       "GooglePlaceID",
		GooglePlaceType:     "GooglePlaceType",
		ReplyMarkup:         &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
		InputMessageContent: &InputTextMessageContent{},
		ThumbURL:            "ThumbURL",
		ThumbWidth:          1,
		ThumbHeight:         2,
	}, i)
}

func TestInlineQueryResultContact_Setters(t *testing.T) {
	i := (&InlineQueryResultContact{}).
		WithID("ID").
		WithPhoneNumber("PhoneNumber").
		WithFirstName("FirstName").
		WithLastName("LastName").
		WithVcard("Vcard").
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}}).
		WithInputMessageContent(&InputTextMessageContent{}).
		WithThumbURL("ThumbURL").
		WithThumbWidth(1).
		WithThumbHeight(2)

	assert.Equal(t, &InlineQueryResultContact{
		ID:                  "ID",
		PhoneNumber:         "PhoneNumber",
		FirstName:           "FirstName",
		LastName:            "LastName",
		Vcard:               "Vcard",
		ReplyMarkup:         &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
		InputMessageContent: &InputTextMessageContent{},
		ThumbURL:            "ThumbURL",
		ThumbWidth:          1,
		ThumbHeight:         2,
	}, i)
}

func TestInlineQueryResultGame_Setters(t *testing.T) {
	i := (&InlineQueryResultGame{}).
		WithID("ID").
		WithGameShortName("GameShortName").
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}})

	assert.Equal(t, &InlineQueryResultGame{
		ID:            "ID",
		GameShortName: "GameShortName",
		ReplyMarkup:   &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
	}, i)
}

func TestInlineQueryResultCachedPhoto_Setters(t *testing.T) {
	i := (&InlineQueryResultCachedPhoto{}).
		WithID("ID").
		WithPhotoFileID("PhotoFileID").
		WithTitle("Title").
		WithDescription("Description").
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}}).
		WithInputMessageContent(&InputTextMessageContent{})

	assert.Equal(t, &InlineQueryResultCachedPhoto{
		ID:                  "ID",
		PhotoFileID:         "PhotoFileID",
		Title:               "Title",
		Description:         "Description",
		Caption:             "Caption",
		ParseMode:           "ParseMode",
		CaptionEntities:     []MessageEntity{{Type: "CaptionEntities"}},
		ReplyMarkup:         &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
		InputMessageContent: &InputTextMessageContent{},
	}, i)
}

func TestInlineQueryResultCachedGif_Setters(t *testing.T) {
	i := (&InlineQueryResultCachedGif{}).
		WithID("ID").
		WithGifFileID("GifFileID").
		WithTitle("Title").
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}}).
		WithInputMessageContent(&InputTextMessageContent{})

	assert.Equal(t, &InlineQueryResultCachedGif{
		ID:                  "ID",
		GifFileID:           "GifFileID",
		Title:               "Title",
		Caption:             "Caption",
		ParseMode:           "ParseMode",
		CaptionEntities:     []MessageEntity{{Type: "CaptionEntities"}},
		ReplyMarkup:         &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
		InputMessageContent: &InputTextMessageContent{},
	}, i)
}

func TestInlineQueryResultCachedMpeg4Gif_Setters(t *testing.T) {
	i := (&InlineQueryResultCachedMpeg4Gif{}).
		WithID("ID").
		WithMpeg4FileID("Mpeg4FileID").
		WithTitle("Title").
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}}).
		WithInputMessageContent(&InputTextMessageContent{})

	assert.Equal(t, &InlineQueryResultCachedMpeg4Gif{
		ID:                  "ID",
		Mpeg4FileID:         "Mpeg4FileID",
		Title:               "Title",
		Caption:             "Caption",
		ParseMode:           "ParseMode",
		CaptionEntities:     []MessageEntity{{Type: "CaptionEntities"}},
		ReplyMarkup:         &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
		InputMessageContent: &InputTextMessageContent{},
	}, i)
}

func TestInlineQueryResultCachedSticker_Setters(t *testing.T) {
	i := (&InlineQueryResultCachedSticker{}).
		WithID("ID").
		WithStickerFileID("StickerFileID").
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}}).
		WithInputMessageContent(&InputTextMessageContent{})

	assert.Equal(t, &InlineQueryResultCachedSticker{
		ID:                  "ID",
		StickerFileID:       "StickerFileID",
		ReplyMarkup:         &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
		InputMessageContent: &InputTextMessageContent{},
	}, i)
}

func TestInlineQueryResultCachedDocument_Setters(t *testing.T) {
	i := (&InlineQueryResultCachedDocument{}).
		WithID("ID").
		WithTitle("Title").
		WithDocumentFileID("DocumentFileID").
		WithDescription("Description").
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}}).
		WithInputMessageContent(&InputTextMessageContent{})

	assert.Equal(t, &InlineQueryResultCachedDocument{
		ID:                  "ID",
		Title:               "Title",
		DocumentFileID:      "DocumentFileID",
		Description:         "Description",
		Caption:             "Caption",
		ParseMode:           "ParseMode",
		CaptionEntities:     []MessageEntity{{Type: "CaptionEntities"}},
		ReplyMarkup:         &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
		InputMessageContent: &InputTextMessageContent{},
	}, i)
}

func TestInlineQueryResultCachedVideo_Setters(t *testing.T) {
	i := (&InlineQueryResultCachedVideo{}).
		WithID("ID").
		WithVideoFileID("VideoFileID").
		WithTitle("Title").
		WithDescription("Description").
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}}).
		WithInputMessageContent(&InputTextMessageContent{})

	assert.Equal(t, &InlineQueryResultCachedVideo{
		ID:                  "ID",
		VideoFileID:         "VideoFileID",
		Title:               "Title",
		Description:         "Description",
		Caption:             "Caption",
		ParseMode:           "ParseMode",
		CaptionEntities:     []MessageEntity{{Type: "CaptionEntities"}},
		ReplyMarkup:         &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
		InputMessageContent: &InputTextMessageContent{},
	}, i)
}

func TestInlineQueryResultCachedVoice_Setters(t *testing.T) {
	i := (&InlineQueryResultCachedVoice{}).
		WithID("ID").
		WithVoiceFileID("VoiceFileID").
		WithTitle("Title").
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}}).
		WithInputMessageContent(&InputTextMessageContent{})

	assert.Equal(t, &InlineQueryResultCachedVoice{
		ID:                  "ID",
		VoiceFileID:         "VoiceFileID",
		Title:               "Title",
		Caption:             "Caption",
		ParseMode:           "ParseMode",
		CaptionEntities:     []MessageEntity{{Type: "CaptionEntities"}},
		ReplyMarkup:         &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
		InputMessageContent: &InputTextMessageContent{},
	}, i)
}

func TestInlineQueryResultCachedAudio_Setters(t *testing.T) {
	i := (&InlineQueryResultCachedAudio{}).
		WithID("ID").
		WithAudioFileID("AudioFileID").
		WithCaption("Caption").
		WithParseMode("ParseMode").
		WithCaptionEntities([]MessageEntity{{Type: "CaptionEntities"}}...).
		WithReplyMarkup(&InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}}).
		WithInputMessageContent(&InputTextMessageContent{})

	assert.Equal(t, &InlineQueryResultCachedAudio{
		ID:                  "ID",
		AudioFileID:         "AudioFileID",
		Caption:             "Caption",
		ParseMode:           "ParseMode",
		CaptionEntities:     []MessageEntity{{Type: "CaptionEntities"}},
		ReplyMarkup:         &InlineKeyboardMarkup{InlineKeyboard: [][]InlineKeyboardButton{{}}},
		InputMessageContent: &InputTextMessageContent{},
	}, i)
}

func TestInputTextMessageContent_Setters(t *testing.T) {
	i := (&InputTextMessageContent{}).
		WithMessageText("MessageText").
		WithParseMode("ParseMode").
		WithEntities([]MessageEntity{{Type: "Entities"}}...).
		WithDisableWebPagePreview()

	assert.Equal(t, &InputTextMessageContent{
		MessageText:           "MessageText",
		ParseMode:             "ParseMode",
		Entities:              []MessageEntity{{Type: "Entities"}},
		DisableWebPagePreview: true,
	}, i)
}

func TestInputLocationMessageContent_Setters(t *testing.T) {
	i := (&InputLocationMessageContent{}).
		WithLivePeriod(1).
		WithHeading(1).
		WithProximityAlertRadius(2)

	assert.Equal(t, &InputLocationMessageContent{
		LivePeriod:           1,
		Heading:              1,
		ProximityAlertRadius: 2,
	}, i)
}

func TestInputVenueMessageContent_Setters(t *testing.T) {
	i := (&InputVenueMessageContent{}).
		WithTitle("Title").
		WithAddress("Address").
		WithFoursquareID("FoursquareID").
		WithFoursquareType("FoursquareType").
		WithGooglePlaceID("GooglePlaceID").
		WithGooglePlaceType("GooglePlaceType")

	assert.Equal(t, &InputVenueMessageContent{
		Title:           "Title",
		Address:         "Address",
		FoursquareID:    "FoursquareID",
		FoursquareType:  "FoursquareType",
		GooglePlaceID:   "GooglePlaceID",
		GooglePlaceType: "GooglePlaceType",
	}, i)
}

func TestInputContactMessageContent_Setters(t *testing.T) {
	i := (&InputContactMessageContent{}).
		WithPhoneNumber("PhoneNumber").
		WithFirstName("FirstName").
		WithLastName("LastName").
		WithVcard("Vcard")

	assert.Equal(t, &InputContactMessageContent{
		PhoneNumber: "PhoneNumber",
		FirstName:   "FirstName",
		LastName:    "LastName",
		Vcard:       "Vcard",
	}, i)
}

func TestInputInvoiceMessageContent_Setters(t *testing.T) {
	i := (&InputInvoiceMessageContent{}).
		WithTitle("Title").
		WithDescription("Description").
		WithPayload("Payload").
		WithProviderToken("ProviderToken").
		WithCurrency("Currency").
		WithPrices([]LabeledPrice{{Label: "Prices"}}...).
		WithMaxTipAmount(1).
		WithSuggestedTipAmounts([]int{2}...).
		WithProviderData("ProviderData").
		WithPhotoURL("PhotoURL").
		WithPhotoSize(3).
		WithPhotoWidth(4).
		WithPhotoHeight(5).
		WithNeedName().
		WithNeedPhoneNumber().
		WithNeedEmail().
		WithNeedShippingAddress().
		WithSendPhoneNumberToProvider().
		WithSendEmailToProvider().
		WithIsFlexible()

	assert.Equal(t, &InputInvoiceMessageContent{
		Title:                     "Title",
		Description:               "Description",
		Payload:                   "Payload",
		ProviderToken:             "ProviderToken",
		Currency:                  "Currency",
		Prices:                    []LabeledPrice{{Label: "Prices"}},
		MaxTipAmount:              1,
		SuggestedTipAmounts:       []int{2},
		ProviderData:              "ProviderData",
		PhotoURL:                  "PhotoURL",
		PhotoSize:                 3,
		PhotoWidth:                4,
		PhotoHeight:               5,
		NeedName:                  true,
		NeedPhoneNumber:           true,
		NeedEmail:                 true,
		NeedShippingAddress:       true,
		SendPhoneNumberToProvider: true,
		SendEmailToProvider:       true,
		IsFlexible:                true,
	}, i)
}
