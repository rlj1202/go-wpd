#pragma once

#include <Windows.h>

#ifdef __cplusplus
extern "C" {
#endif

	typedef WCHAR *PnPDeviceID;
	
	extern const CLSID CLSID_PortableDeviceManager;

	extern const IID IID_IPortableDeviceValues;
	extern const IID IID_IPortableDeviceDataStream;

	extern const PROPERTYKEY WPD_CLIENT_NAME;
	extern const PROPERTYKEY WPD_CLIENT_MAJOR_VERSION;
	extern const PROPERTYKEY WPD_CLIENT_MINOR_VERSION;
	extern const PROPERTYKEY WPD_CLIENT_REVISION;
	extern const PROPERTYKEY WPD_CLIENT_SECURITY_QUALITY_OF_SERVICE;
	extern const PROPERTYKEY WPD_CLIENT_DESIRED_ACCESS;

	extern const PROPERTYKEY WPD_OBJECT_PARENT_ID;
	extern const PROPERTYKEY WPD_OBJECT_NAME;
	extern const PROPERTYKEY WPD_OBJECT_PERSISTENT_UNIQUE_ID;
	extern const PROPERTYKEY WPD_OBJECT_FORMAT;
	extern const PROPERTYKEY WPD_OBJECT_CONTENT_TYPE;
	extern const PROPERTYKEY WPD_OBJECT_SIZE;
	extern const PROPERTYKEY WPD_OBJECT_ORIGINAL_FILE_NAME;

	extern const PROPERTYKEY WPD_PROPERTY_ATTRIBUTE_FORM;
	extern const PROPERTYKEY WPD_PROPERTY_ATTRIBUTE_CAN_READ;
	extern const PROPERTYKEY WPD_PROPERTY_ATTRIBUTE_CAN_WRITE;
	extern const PROPERTYKEY WPD_PROPERTY_ATTRIBUTE_CAN_DELETE;
	extern const PROPERTYKEY WPD_PROPERTY_ATTRIBUTE_DEFAULT_VALUE;
	extern const PROPERTYKEY WPD_PROPERTY_ATTRIBUTE_FAST_PROPERTY;
	extern const PROPERTYKEY WPD_PROPERTY_ATTRIBUTE_RANGE_MIN;
	extern const PROPERTYKEY WPD_PROPERTY_ATTRIBUTE_RANGE_MAX;
	extern const PROPERTYKEY WPD_PROPERTY_ATTRIBUTE_RANGE_STEP;
	extern const PROPERTYKEY WPD_PROPERTY_ATTRIBUTE_ENUMERATION_ELEMENTS;
	extern const PROPERTYKEY WPD_PROPERTY_ATTRIBUTE_REGULAR_EXPRESSION;
	extern const PROPERTYKEY WPD_PROPERTY_ATTRIBUTE_MAX_SIZE;

	extern const GUID WPD_CONTENT_TYPE_FUNCTIONAL_OBJECT;
	extern const GUID WPD_CONTENT_TYPE_FOLDER;
	extern const GUID WPD_CONTENT_TYPE_IMAGE;
	extern const GUID WPD_CONTENT_TYPE_DOCUMENT;
	extern const GUID WPD_CONTENT_TYPE_CONTACT;
	extern const GUID WPD_CONTENT_TYPE_CONTACT_GROUP;
	extern const GUID WPD_CONTENT_TYPE_AUDIO;
	extern const GUID WPD_CONTENT_TYPE_VIDEO;
	extern const GUID WPD_CONTENT_TYPE_TELEVISION;
	extern const GUID WPD_CONTENT_TYPE_PLAYLIST;
	extern const GUID WPD_CONTENT_TYPE_MIXED_CONTENT_ALBUM;
	extern const GUID WPD_CONTENT_TYPE_AUDIO_ALBUM;
	extern const GUID WPD_CONTENT_TYPE_IMAGE_ALBUM;
	extern const GUID WPD_CONTENT_TYPE_VIDEO_ALBUM;
	extern const GUID WPD_CONTENT_TYPE_MEMO;
	extern const GUID WPD_CONTENT_TYPE_EMAIL;
	extern const GUID WPD_CONTENT_TYPE_APPOINTMENT;
	extern const GUID WPD_CONTENT_TYPE_TASK;
	extern const GUID WPD_CONTENT_TYPE_PROGRAM;
	extern const GUID WPD_CONTENT_TYPE_GENERIC_FILE;
	extern const GUID WPD_CONTENT_TYPE_CALENDAR;
	extern const GUID WPD_CONTENT_TYPE_GENERIC_MESSAGE;
	extern const GUID WPD_CONTENT_TYPE_NETWORK_ASSOCIATION;
	extern const GUID WPD_CONTENT_TYPE_CERTIFICATE;
	extern const GUID WPD_CONTENT_TYPE_WIRELESS_PROFILE;
	extern const GUID WPD_CONTENT_TYPE_MEDIA_CAST;
	extern const GUID WPD_CONTENT_TYPE_SECTION;
	extern const GUID WPD_CONTENT_TYPE_UNSPECIFIED;
	extern const GUID WPD_CONTENT_TYPE_ALL;

	extern const GUID WPD_OBJECT_FORMAT_EXIF;
	extern const GUID WPD_OBJECT_FORMAT_WMA;
	extern const GUID WPD_OBJECT_FORMAT_VCARD2;

	typedef struct IPortableDevice IPortableDevice;
	typedef struct IPortableDeviceValues IPortableDeviceValues;
	typedef struct IPortableDeviceManager IPortableDeviceManager;
	typedef struct IPortableDeviceContent IPortableDeviceContent;
	typedef struct IPortableDeviceKeyCollection IPortableDeviceKeyCollection;
	typedef struct IPortableDeviceProperties IPortableDeviceProperties;
	typedef struct IPortableDeviceDataStream IPortableDeviceDataStream;
	typedef struct IPortableDeviceCapabilities IPortableDeviceCapabilities;
	typedef struct IPortableDevicePropVariantCollection IPortableDevicePropVariantCollection;
	typedef struct IPortableDeviceEventCallback IPortableDeviceEventCallback;
	typedef struct IStream IStream;
	typedef struct ISequentialStream ISequentialStream;
	typedef struct IPropertyStore IPropertyStore;
	typedef struct IUnknown IUnknown;
	
	typedef struct IEnumPortableDeviceObjectIDs IEnumPortableDeviceObjectIDs;
	
	HRESULT createPortableDevice(IPortableDevice **ppPortableDevice);
	HRESULT createPortableDeviceValues(IPortableDeviceValues **ppPortableDeviceValues);
	HRESULT createPortableDeviceManager(IPortableDeviceManager **ppPortableDeviceManager);
	HRESULT createPortableDeviceKeyCollection(IPortableDeviceKeyCollection **ppPortableDeviceKeyCollection);

	HRESULT portableDevice_Advise(IPortableDevice *pPortableDevice, DWORD flags, IPortableDeviceEventCallback *pCallback, IPortableDeviceValues *pParameters, PWSTR *ppCookie);
	HRESULT portableDevice_Cancel(IPortableDevice *pPortableDevice);
	HRESULT portableDevice_Capabilities(IPortableDevice *pPortableDevice, IPortableDeviceCapabilities **ppCapabilities);
	HRESULT portableDevice_Close(IPortableDevice *pPortableDevice);
	HRESULT portableDevice_Content(IPortableDevice *pPortableDevice, IPortableDeviceContent **ppPortableDeviceContent);
	HRESULT portableDevice_GetPnPDeviceID(IPortableDevice *pPortableDevice, PnPDeviceID *pPnPDeviceID);
	HRESULT portableDevice_Open(IPortableDevice *pPortableDevice, PnPDeviceID pnpDeviceID, IPortableDeviceValues *pClientInfo);
	HRESULT portableDevice_SendCommand(IPortableDevice *pPortableDevice, DWORD flags, IPortableDeviceValues *pValues, IPortableDeviceValues **ppResults);
	HRESULT portableDevice_Unadvise(IPortableDevice *pPortableDevice, PWSTR pCookie);
	HRESULT portableDevice_Release(IPortableDevice *pPortableDevice);

	HRESULT portableDeviceValues_GetBoolValue(IPortableDeviceValues *pPortableDeviceValues, const PROPERTYKEY *key, BOOL *value);
	HRESULT portableDeviceValues_GetStringValue(IPortableDeviceValues *pPortableDeviceValues, const PROPERTYKEY *key, PWSTR *value, DWORD *cValue);
	HRESULT portableDeviceValues_GetUnsignedIntegerValue(IPortableDeviceValues *pPortableDeviceValues, const PROPERTYKEY *key, ULONG *value);
	HRESULT portableDeviceValues_SetGuidValue(IPortableDeviceValues *pPortableDeviceValues, const PROPERTYKEY *key, GUID *pGuid);
	HRESULT portableDeviceValues_SetStringValue(IPortableDeviceValues *pPortableDeviceValues, const PROPERTYKEY *key, LPCWSTR value);
	HRESULT portableDeviceValues_SetUnsignedIntegerValue(IPortableDeviceValues *pPortableDeviceValues, const PROPERTYKEY *key, const ULONG value);
	HRESULT portableDeviceValues_SetUnsignedLargeIntegerValue(IPortableDeviceValues *pPortableDeviceValues, const PROPERTYKEY *key, const ULONGLONG value);
	HRESULT portableDeviceValues_Release(IPortableDeviceValues *pPortableDeviceValues);

	HRESULT portableDeviceManager_GetDevices(IPortableDeviceManager *pPortableDeviceManager, PnPDeviceID **pPnPDeviceIDs, DWORD *cPnPDeviceIDs);
	HRESULT portableDeviceManager_GetDeviceFriendlyName(IPortableDeviceManager *pPortableDeviceManager, PnPDeviceID pnpDeviceID, PWSTR *pFriendlyName, DWORD *cFriendlyName);
	HRESULT portableDeviceManager_GetDeviceManufacturer(IPortableDeviceManager *pPortableDeviceManager, PnPDeviceID pnpDeviceID, PWSTR *pManufacturer, DWORD *cManufacturer);
	HRESULT portableDeviceManager_GetDeviceDescription(IPortableDeviceManager *pPortableDeviceManager, PnPDeviceID pnpDeviceID, PWSTR *pDescription, DWORD *cDescription);

	HRESULT portableDeviceContent_CreateObjectWithPropertiesAndData(IPortableDeviceContent *pPortableDeviceContent, IPortableDeviceValues *pValues, IStream **ppData, DWORD *pOptimalWriteBufferSize, PWSTR *ppCookie);
	HRESULT portableDeviceContent_EnumObjects(IPortableDeviceContent *pPortableDeviceContent, DWORD flags, PWSTR parentObjectID, IPortableDeviceValues *pFilter, IEnumPortableDeviceObjectIDs **ppEnum);
	HRESULT portableDeviceContent_Properties(IPortableDeviceContent *pPortableDeviceContent, IPortableDeviceProperties **ppPortableDeviceProperties);

	HRESULT portableDeviceKeyCollection_Add(IPortableDeviceKeyCollection *pPortableDeviceKeyCollection, const PROPERTYKEY *key);
	
	HRESULT portableDeviceProperties_GetValues(IPortableDeviceProperties *pPortableDeviceProperties, PWSTR pObjectID, IPortableDeviceKeyCollection *pKeys, IPortableDeviceValues **ppValues);
	HRESULT portableDeviceProperties_GetPropertyAttributes(IPortableDeviceProperties *pPortableDeviceProperties, PWSTR pObjectID, const PROPERTYKEY *key, IPortableDeviceValues **ppAttributes);
	HRESULT portableDeviceProperties_SetValues(IPortableDeviceProperties *pPortableDeviceProperties, PWSTR pObjectID, IPortableDeviceValues *pValues, IPortableDeviceValues **ppResults);

	HRESULT portableDeviceDataStream_Commit(IPortableDeviceDataStream *pPortableDeviceDataStream, DWORD dataFlags);
	HRESULT portableDeviceDataStream_GetObjectID(IPortableDeviceDataStream *pPortableDeviceDataStream, PWSTR *pObjectID);

	HRESULT enumPortableDeviceObjectIDs_Next(IEnumPortableDeviceObjectIDs *pEnumObjectIDs, ULONG cObjects, PWSTR *pObjIDs, DWORD *pcObjIDs, ULONG *pcPetched);

	HRESULT stream_Commit(IStream *pStream, DWORD dataFlag);
	HRESULT stream_Stat(IStream *pStream, STATSTG *pStatstg, DWORD statFlags);

	HRESULT sequentialStream_Read(ISequentialStream *pSequentialStream, LPVOID pBuffer, ULONG cb, ULONG *pcbRead);
	HRESULT sequentialStream_Write(ISequentialStream *pSequentialStream, LPVOID pBuffer, ULONG cb, ULONG *pcbWritten);
	
	HRESULT unknown_QueryInterface(IUnknown *pUnknown, const IID *piid, LPVOID *ppvObject);

	void test();

#ifdef __cplusplus
}
#endif
