# Paging

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Previous** | **string** |  | 
**Next** | **string** |  | 

## Methods

### NewPaging

`func NewPaging(previous string, next string, ) *Paging`

NewPaging instantiates a new Paging object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagingWithDefaults

`func NewPagingWithDefaults() *Paging`

NewPagingWithDefaults instantiates a new Paging object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrevious

`func (o *Paging) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *Paging) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *Paging) SetPrevious(v string)`

SetPrevious sets Previous field to given value.


### GetNext

`func (o *Paging) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *Paging) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *Paging) SetNext(v string)`

SetNext sets Next field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


