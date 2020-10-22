//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
package model

import (
    "errors"
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/spi"
    "strconv"
)

// Constant values.
const BACnetErrorReadProperty_ERRORCLASSHEADER uint8 = 0x12
const BACnetErrorReadProperty_ERRORCODEHEADER uint8 = 0x12

// The data-structure of this message
type BACnetErrorReadProperty struct {
    ErrorClassLength uint8
    ErrorClass []int8
    ErrorCodeLength uint8
    ErrorCode []int8
    BACnetError
}

// The corresponding interface
type IBACnetErrorReadProperty interface {
    IBACnetError
    Serialize(io spi.WriteBuffer) error
}

// Accessors for discriminator values.
func (m BACnetErrorReadProperty) ServiceChoice() uint8 {
    return 0x0C
}

func (m BACnetErrorReadProperty) initialize() spi.Message {
    return m
}

func NewBACnetErrorReadProperty(errorClassLength uint8, errorClass []int8, errorCodeLength uint8, errorCode []int8) BACnetErrorInitializer {
    return &BACnetErrorReadProperty{ErrorClassLength: errorClassLength, ErrorClass: errorClass, ErrorCodeLength: errorCodeLength, ErrorCode: errorCode}
}

func CastIBACnetErrorReadProperty(structType interface{}) IBACnetErrorReadProperty {
    castFunc := func(typ interface{}) IBACnetErrorReadProperty {
        if iBACnetErrorReadProperty, ok := typ.(IBACnetErrorReadProperty); ok {
            return iBACnetErrorReadProperty
        }
        return nil
    }
    return castFunc(structType)
}

func CastBACnetErrorReadProperty(structType interface{}) BACnetErrorReadProperty {
    castFunc := func(typ interface{}) BACnetErrorReadProperty {
        if sBACnetErrorReadProperty, ok := typ.(BACnetErrorReadProperty); ok {
            return sBACnetErrorReadProperty
        }
        if sBACnetErrorReadProperty, ok := typ.(*BACnetErrorReadProperty); ok {
            return *sBACnetErrorReadProperty
        }
        return BACnetErrorReadProperty{}
    }
    return castFunc(structType)
}

func (m BACnetErrorReadProperty) LengthInBits() uint16 {
    var lengthInBits uint16 = m.BACnetError.LengthInBits()

    // Const Field (errorClassHeader)
    lengthInBits += 5

    // Simple field (errorClassLength)
    lengthInBits += 3

    // Array field
    if len(m.ErrorClass) > 0 {
        lengthInBits += 8 * uint16(len(m.ErrorClass))
    }

    // Const Field (errorCodeHeader)
    lengthInBits += 5

    // Simple field (errorCodeLength)
    lengthInBits += 3

    // Array field
    if len(m.ErrorCode) > 0 {
        lengthInBits += 8 * uint16(len(m.ErrorCode))
    }

    return lengthInBits
}

func (m BACnetErrorReadProperty) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func BACnetErrorReadPropertyParse(io *spi.ReadBuffer) (BACnetErrorInitializer, error) {

    // Const Field (errorClassHeader)
    errorClassHeader, _errorClassHeaderErr := io.ReadUint8(5)
    if _errorClassHeaderErr != nil {
        return nil, errors.New("Error parsing 'errorClassHeader' field " + _errorClassHeaderErr.Error())
    }
    if errorClassHeader != BACnetErrorReadProperty_ERRORCLASSHEADER {
        return nil, errors.New("Expected constant value " + strconv.Itoa(int(BACnetErrorReadProperty_ERRORCLASSHEADER)) + " but got " + strconv.Itoa(int(errorClassHeader)))
    }

    // Simple Field (errorClassLength)
    errorClassLength, _errorClassLengthErr := io.ReadUint8(3)
    if _errorClassLengthErr != nil {
        return nil, errors.New("Error parsing 'errorClassLength' field " + _errorClassLengthErr.Error())
    }

    // Array field (errorClass)
    // Count array
    errorClass := make([]int8, errorClassLength)
    for curItem := uint16(0); curItem < uint16(errorClassLength); curItem++ {

        _item, _err := io.ReadInt8(8)
        if _err != nil {
            return nil, errors.New("Error parsing 'errorClass' field " + _err.Error())
        }
        errorClass[curItem] = _item
    }

    // Const Field (errorCodeHeader)
    errorCodeHeader, _errorCodeHeaderErr := io.ReadUint8(5)
    if _errorCodeHeaderErr != nil {
        return nil, errors.New("Error parsing 'errorCodeHeader' field " + _errorCodeHeaderErr.Error())
    }
    if errorCodeHeader != BACnetErrorReadProperty_ERRORCODEHEADER {
        return nil, errors.New("Expected constant value " + strconv.Itoa(int(BACnetErrorReadProperty_ERRORCODEHEADER)) + " but got " + strconv.Itoa(int(errorCodeHeader)))
    }

    // Simple Field (errorCodeLength)
    errorCodeLength, _errorCodeLengthErr := io.ReadUint8(3)
    if _errorCodeLengthErr != nil {
        return nil, errors.New("Error parsing 'errorCodeLength' field " + _errorCodeLengthErr.Error())
    }

    // Array field (errorCode)
    // Count array
    errorCode := make([]int8, errorCodeLength)
    for curItem := uint16(0); curItem < uint16(errorCodeLength); curItem++ {

        _item, _err := io.ReadInt8(8)
        if _err != nil {
            return nil, errors.New("Error parsing 'errorCode' field " + _err.Error())
        }
        errorCode[curItem] = _item
    }

    // Create the instance
    return NewBACnetErrorReadProperty(errorClassLength, errorClass, errorCodeLength, errorCode), nil
}

func (m BACnetErrorReadProperty) Serialize(io spi.WriteBuffer) error {
    ser := func() error {

    // Const Field (errorClassHeader)
    _errorClassHeaderErr := io.WriteUint8(5, 0x12)
    if _errorClassHeaderErr != nil {
        return errors.New("Error serializing 'errorClassHeader' field " + _errorClassHeaderErr.Error())
    }

    // Simple Field (errorClassLength)
    errorClassLength := uint8(m.ErrorClassLength)
    _errorClassLengthErr := io.WriteUint8(3, (errorClassLength))
    if _errorClassLengthErr != nil {
        return errors.New("Error serializing 'errorClassLength' field " + _errorClassLengthErr.Error())
    }

    // Array Field (errorClass)
    if m.ErrorClass != nil {
        for _, _element := range m.ErrorClass {
            _elementErr := io.WriteInt8(8, _element)
            if _elementErr != nil {
                return errors.New("Error serializing 'errorClass' field " + _elementErr.Error())
            }
        }
    }

    // Const Field (errorCodeHeader)
    _errorCodeHeaderErr := io.WriteUint8(5, 0x12)
    if _errorCodeHeaderErr != nil {
        return errors.New("Error serializing 'errorCodeHeader' field " + _errorCodeHeaderErr.Error())
    }

    // Simple Field (errorCodeLength)
    errorCodeLength := uint8(m.ErrorCodeLength)
    _errorCodeLengthErr := io.WriteUint8(3, (errorCodeLength))
    if _errorCodeLengthErr != nil {
        return errors.New("Error serializing 'errorCodeLength' field " + _errorCodeLengthErr.Error())
    }

    // Array Field (errorCode)
    if m.ErrorCode != nil {
        for _, _element := range m.ErrorCode {
            _elementErr := io.WriteInt8(8, _element)
            if _elementErr != nil {
                return errors.New("Error serializing 'errorCode' field " + _elementErr.Error())
            }
        }
    }

        return nil
    }
    return BACnetErrorSerialize(io, m.BACnetError, CastIBACnetError(m), ser)
}