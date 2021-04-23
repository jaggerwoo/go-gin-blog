package qrcode

import (
    "image/jpeg"

    "github.com/boombuler/barcode"
    "github.com/boombuler/barcode/qr"
)

type QrCode struct {
    URL string
    Width int
    Height int
    Ext string
    Level  qr.ErrorCorrectionLevel
    Mode   qr.Encoding
}