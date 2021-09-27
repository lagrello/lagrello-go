# TokenList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Token**](Token.md) |  | 
**Paging** | [**Paging**](Paging.md) |  | 

## Methods

### NewTokenList

`func NewTokenList(data []Token, paging Paging, ) *TokenList`

NewTokenList instantiates a new TokenList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenListWithDefaults

`func NewTokenListWithDefaults() *TokenList`

NewTokenListWithDefaults instantiates a new TokenList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TokenList) GetData() []Token`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TokenList) GetDataOk() (*[]Token, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TokenList) SetData(v []Token)`

SetData sets Data field to given value.


### GetPaging

`func (o *TokenList) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *TokenList) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *TokenList) SetPaging(v Paging)`

SetPaging sets Paging field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


