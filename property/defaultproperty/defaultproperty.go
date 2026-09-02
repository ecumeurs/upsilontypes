package defaultproperty

import (
	"fmt"

	"github.com/ecumeurs/upsilontools/tools"
	"github.com/ecumeurs/upsilontypes/property"
)

type DefaultIntProperty struct {
	value               int
	name                string
	minInformationLevel property.InformationLevel
	propertyType        property.PropertyType
}

// MakeIntProperty builds a DefaultIntProperty for the given property key.
func MakeIntProperty(name property.Key, value int, minInformationLevel property.InformationLevel, t property.PropertyType) *DefaultIntProperty {
	return &DefaultIntProperty{
		name:                name.String(),
		value:               value,
		minInformationLevel: minInformationLevel,
		propertyType:        t,
	}
}

// implements IntProperty
func (d DefaultIntProperty) Name(i property.InformationLevel) string {
	if i >= d.minInformationLevel {
		return d.name
	}
	return ""
}

func (d DefaultIntProperty) UserFriendlyGet(i property.InformationLevel) interface{} {
	if i >= d.minInformationLevel {
		return d.value
	}
	return nil
}

func (d DefaultIntProperty) Get() interface{} {
	return d.value
}

func (d *DefaultIntProperty) Set(p interface{}) {
	d.value = p.(int)
}

func (d DefaultIntProperty) Increase() {
	// do nothing
}

func (d DefaultIntProperty) GetType() property.PropertyType {
	return d.propertyType
}

func (d DefaultIntProperty) I() int {
	return d.value
}

func (d *DefaultIntProperty) SetI(i int) {
	d.value = i
}

func (d DefaultIntProperty) ApplyBuff(p property.Property) property.Property {
	res := d.Duplicate()
	res.Set(d.Get().(int) + p.Get().(int))
	return res
}

func (d DefaultIntProperty) UnapplyBuff(p property.Property) property.Property {
	res := d.Duplicate()
	res.Set(d.Get().(int) - p.Get().(int))
	return res
}

func (d DefaultIntProperty) Duplicate() property.Property {
	return &DefaultIntProperty{
		value:               d.value,
		name:                d.name,
		minInformationLevel: d.minInformationLevel,
		propertyType:        d.propertyType,
	}
}

type DefaultIntCounterProperty struct {
	Value    int
	MaxValue int

	name                string
	minInformationLevel property.InformationLevel
	propertyType        property.PropertyType
}

// MakeIntCounterProperty builds a DefaultIntCounterProperty for the given property key.
func MakeIntCounterProperty(name property.Key, value, maxvalue int, minInformationLevel property.InformationLevel, t property.PropertyType) *DefaultIntCounterProperty {
	nname := name.String()
	if nname == "" {
		panic(fmt.Sprintf("defaultproperty: MakeIntCounterProperty called with an empty property key (value=%d, maxvalue=%d); a nil Property must never enter the property map", value, maxvalue))
	}

	return &DefaultIntCounterProperty{
		name:                nname,
		Value:               value,
		MaxValue:            maxvalue,
		minInformationLevel: minInformationLevel,
		propertyType:        t,
	}
}

func (d DefaultIntCounterProperty) GetValue() int {
	return d.Value
}

func (d DefaultIntCounterProperty) GetMaxValue() int {
	return d.MaxValue
}

func (d *DefaultIntCounterProperty) SetValue(i int) {
	d.Value = i
}
func (d *DefaultIntCounterProperty) SetMaxValue(i int) {
	d.MaxValue = i
}

// implements IntProperty
func (d DefaultIntCounterProperty) Name(i property.InformationLevel) string {
	if i >= d.minInformationLevel {
		return d.name
	}
	return ""
}

func (d DefaultIntCounterProperty) UserFriendlyGet(i property.InformationLevel) interface{} {
	if i >= d.minInformationLevel {
		return d.Value
	}
	return nil
}

func (d DefaultIntCounterProperty) Get() interface{} {
	return d.Value
}

func (d *DefaultIntCounterProperty) Set(p interface{}) {
	d.Value = p.(int)
}

func (d DefaultIntCounterProperty) Increase() {
	// do nothing
}

func (d DefaultIntCounterProperty) GetType() property.PropertyType {
	return d.propertyType
}

func (d DefaultIntCounterProperty) I() int {
	return d.Value
}

func (d *DefaultIntCounterProperty) SetI(i int) {
	d.Value = i
}

func (d DefaultIntCounterProperty) ApplyBuff(p property.Property) property.Property {
	res := d.Duplicate().(*DefaultIntCounterProperty)
	res.Value = d.Value + p.(*DefaultIntCounterProperty).Value
	res.MaxValue = d.MaxValue + p.(*DefaultIntCounterProperty).MaxValue
	return res
}

func (d DefaultIntCounterProperty) UnapplyBuff(p property.Property) property.Property {
	res := d.Duplicate().(*DefaultIntCounterProperty)

	res.MaxValue = d.MaxValue - p.(*DefaultIntCounterProperty).MaxValue
	res.Value = tools.Max(d.Value-p.(*DefaultIntCounterProperty).Value, res.MaxValue)
	return res
}

func (d DefaultIntCounterProperty) Duplicate() property.Property {
	return &DefaultIntCounterProperty{
		Value:               d.Value,
		MaxValue:            d.MaxValue,
		name:                d.name,
		minInformationLevel: d.minInformationLevel,
		propertyType:        d.propertyType,
	}
}

type DefaultFloatProperty struct {
	value               float64
	name                string
	minInformationLevel property.InformationLevel
	propertyType        property.PropertyType
}

// MakeFloatProperty builds a DefaultFloatProperty for the given property key.
func MakeFloatProperty(name property.Key, value float64, minInformationLevel property.InformationLevel, pt property.PropertyType) *DefaultFloatProperty {
	return &DefaultFloatProperty{
		name:                name.String(),
		value:               value,
		minInformationLevel: minInformationLevel,
		propertyType:        pt,
	}
}

// implements IntProperty
func (d DefaultFloatProperty) Name(i property.InformationLevel) string {
	return d.name
}

func (d DefaultFloatProperty) UserFriendlyGet(i property.InformationLevel) interface{} {
	if i >= d.minInformationLevel {
		return d.value
	}
	return nil
}

func (d DefaultFloatProperty) Get() interface{} {
	return d.value
}

func (d *DefaultFloatProperty) Set(p interface{}) {
	d.value = p.(float64)
}

func (d DefaultFloatProperty) Increase() {
	// do nothing
}

func (d DefaultFloatProperty) GetType() property.PropertyType {
	return d.propertyType
}

func (d DefaultFloatProperty) I() float64 {
	return d.value
}

func (d *DefaultFloatProperty) SetI(f float64) {
	d.value = f
}

func (d DefaultFloatProperty) Duplicate() property.Property {
	return &DefaultFloatProperty{
		value:               d.value,
		name:                d.name,
		minInformationLevel: d.minInformationLevel,
		propertyType:        d.propertyType,
	}
}

func (d DefaultFloatProperty) ApplyBuff(p property.Property) property.Property {
	res := d.Duplicate()
	res.Set(d.Get().(float64) + p.Get().(float64))
	return res
}
func (d DefaultFloatProperty) UnapplyBuff(p property.Property) property.Property {
	res := d.Duplicate()
	res.Set(d.Get().(float64) - p.Get().(float64))
	return res
}

// Bool default value are essentially flags ...

type DefaultBoolProperty struct {
	value               bool
	name                string
	minInformationLevel property.InformationLevel
	propertyType        property.PropertyType
}

// MakeBoolProperty builds a DefaultBoolProperty for the given property key.
func MakeBoolProperty(name property.Key, value bool, minInformationLevel property.InformationLevel, pt property.PropertyType) *DefaultBoolProperty {
	return &DefaultBoolProperty{
		name:                name.String(),
		value:               value,
		minInformationLevel: minInformationLevel,
		propertyType:        pt,
	}
}

// implements IntProperty
func (d DefaultBoolProperty) Name(i property.InformationLevel) string {
	return d.name
}

func (d DefaultBoolProperty) UserFriendlyGet(i property.InformationLevel) interface{} {
	if i >= d.minInformationLevel {
		return d.value
	}
	return nil
}

func (d DefaultBoolProperty) Get() interface{} {
	return d.value
}

func (d *DefaultBoolProperty) Set(p interface{}) {
	d.value = p.(bool)
}

func (d DefaultBoolProperty) Increase() {
	// do nothing
}

func (d DefaultBoolProperty) GetType() property.PropertyType {
	return d.propertyType
}

func (d DefaultBoolProperty) B() bool {
	return d.value
}

func (d *DefaultBoolProperty) SetB(f bool) {
	d.value = f
}

func (d DefaultBoolProperty) Duplicate() property.Property {
	return &DefaultBoolProperty{
		value:               d.value,
		name:                d.name,
		minInformationLevel: d.minInformationLevel,
		propertyType:        d.propertyType,
	}
}

func (d DefaultBoolProperty) ApplyBuff(p property.Property) property.Property {
	res := d.Duplicate()
	res.Set(d.Get().(bool) && p.Get().(bool))
	return res
}

func (d DefaultBoolProperty) UnapplyBuff(p property.Property) property.Property {
	res := d.Duplicate()
	res.Set(!p.Get().(bool))
	return res
}

type DefaultStringProperty struct {
	value               string
	name                string
	minInformationLevel property.InformationLevel
	propertyType        property.PropertyType
	AllowedValues       []string
}

// MakeStringProperty builds a DefaultStringProperty for the given property key.
func MakeStringProperty(name property.Key, value string, minInformationLevel property.InformationLevel, pt property.PropertyType) *DefaultStringProperty {
	return MakeValidatedStringProperty(name, value, minInformationLevel, pt, nil)
}

// MakeValidatedStringProperty builds a DefaultStringProperty for the given property
// key, restricting Set/SetS to the allowed value set.
func MakeValidatedStringProperty(name property.Key, value string, minInformationLevel property.InformationLevel, pt property.PropertyType, allowed []string) *DefaultStringProperty {
	nname := name.String()
	if nname == "" {
		panic(fmt.Sprintf("defaultproperty: MakeValidatedStringProperty called with an empty property key (value=%q); a nil Property must never enter the property map", value))
	}

	return &DefaultStringProperty{
		name:                nname,
		value:               value,
		minInformationLevel: minInformationLevel,
		propertyType:        pt,
		AllowedValues:       allowed,
	}
}

func (d DefaultStringProperty) Name(i property.InformationLevel) string {
	if i >= d.minInformationLevel {
		return d.name
	}
	return ""
}

func (d DefaultStringProperty) UserFriendlyGet(i property.InformationLevel) interface{} {
	if i >= d.minInformationLevel {
		return d.value
	}
	return nil
}

func (d DefaultStringProperty) Get() interface{} {
	return d.value
}

func (d DefaultStringProperty) S() string {
	return d.value
}

func (d *DefaultStringProperty) Set(p interface{}) {
	d.SetS(p.(string))
}

func (d *DefaultStringProperty) SetS(s string) {
	if len(d.AllowedValues) > 0 {
		found := false
		for _, v := range d.AllowedValues {
			if v == s {
				found = true
				break
			}
		}
		if !found {
			return
		}
	}
	d.value = s
}

func (d DefaultStringProperty) Increase() {
	// do nothing
}

func (d DefaultStringProperty) GetType() property.PropertyType {
	return d.propertyType
}

func (d DefaultStringProperty) Duplicate() property.Property {
	return &DefaultStringProperty{
		value:               d.value,
		name:                d.name,
		minInformationLevel: d.minInformationLevel,
		propertyType:        d.propertyType,
		AllowedValues:       d.AllowedValues,
	}
}

func (d DefaultStringProperty) ApplyBuff(p property.Property) property.Property {
	return d.Duplicate()
}

func (d DefaultStringProperty) UnapplyBuff(p property.Property) property.Property {
	return d.Duplicate()
}
