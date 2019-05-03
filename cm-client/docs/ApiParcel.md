# ApiParcel

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | **string** | The name of the product, e.g. CDH, Impala | [optional] [default to null]
**Version** | **string** | The version of the product, e.g. 1.1.0, 2.3.0. | [optional] [default to null]
**Stage** | **string** | Returns the current stage of the parcel. &lt;p&gt; There are a number of stages a parcel can be in. There are two types of stages - stable and transient. A parcel is in a transient stage when it is transitioning between two stable stages. The stages are listed below with some additional information.  &lt;ul&gt; &lt;li&gt;&lt;b&gt;AVAILABLE_REMOTELY&lt;/b&gt;: Stable stage - the parcel can be downloaded to the server.&lt;/li&gt; &lt;li&gt;&lt;b&gt;DOWNLOADING&lt;/b&gt;: Transient stage - the parcel is in the process of being downloaded to the server.&lt;/li&gt; &lt;li&gt;&lt;b&gt;DOWNLOADED&lt;/b&gt;: Stable stage - the parcel is downloaded and ready to be distributed or removed from the server.&lt;/li&gt; &lt;li&gt;&lt;b&gt;DISTRIBUTING&lt;/b&gt;: Transient stage - the parcel is being sent to all the hosts in the cluster.&lt;/li&gt; &lt;li&gt;&lt;b&gt;DISTRIBUTED&lt;/b&gt;: Stable stage - the parcel is on all the hosts in the cluster. The parcel can now be activated, or removed from all the hosts.&lt;/li&gt; &lt;li&gt;&lt;b&gt;UNDISTRIBUTING&lt;/b&gt;: Transient stage - the parcel is being removed from all the hosts in the cluster&gt;&lt;/li&gt; &lt;li&gt;&lt;b&gt;ACTIVATING&lt;/b&gt;: Transient stage - the parcel is being activated on the hosts in the cluster. &lt;i&gt;New in API v7&lt;/i&gt;&lt;/li&gt; &lt;li&gt;&lt;b&gt;ACTIVATED&lt;/b&gt;: Steady stage - the parcel is set to active on every host in the cluster. If desired, a parcel can be deactivated from this stage.&lt;/li&gt; &lt;/ul&gt; | [optional] [default to null]
**State** | [***ApiParcelState**](ApiParcelState.md) | The state of the parcel. This shows the progress of state transitions and if there were any errors. | [optional] [default to null]
**ClusterRef** | [***ApiClusterRef**](ApiClusterRef.md) | Readonly. A reference to the enclosing cluster. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


