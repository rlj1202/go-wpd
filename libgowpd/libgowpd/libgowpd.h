#pragma once

#include <Windows.h>

#ifdef __cplusplus
extern "C" {
#endif

	typedef WCHAR *PnPDeviceID;

	extern const PROPERTYKEY WPD_CLIENT_NAME;
	extern const PROPERTYKEY WPD_CLIENT_MAJOR_VERSION;
	extern const PROPERTYKEY WPD_CLIENT_MINOR_VERSION;
	extern const PROPERTYKEY WPD_CLIENT_REVISION;
	extern const PROPERTYKEY WPD_CLIENT_SECURITY_QUALITY_OF_SERVICE;
	extern const PROPERTYKEY WPD_CLIENT_DESIRED_ACCESS;

	typedef struct IPortableDevice IPortableDevice;
	typedef struct IPortableDeviceValues IPortableDeviceValues;
	typedef struct IPortableDeviceManager IPortableDeviceManager;
	typedef struct IPortableDeviceContent IPortableDeviceContent;

	typedef struct IEnumPortableDeviceObjectIDs IEnumPortableDeviceObjectIDs;
	
	HRESULT createPortableDevice(IPortableDevice **ppPortableDevice);
	HRESULT createPortableDeviceValues(IPortableDeviceValues **ppPortableDeviceValues);
	HRESULT createPortableDeviceManager(IPortableDeviceManager **ppPortableDeviceManager);

	HRESULT portableDevice_Open(IPortableDevice *pPortableDevice, PnPDeviceID pnpDeviceID, IPortableDeviceValues *pClientInfo);
	HRESULT portableDevice_Content(IPortableDevice *pPortableDevice, IPortableDeviceContent **ppPortableDeviceContent);
	HRESULT portableDevice_Release(IPortableDevice *pPortableDevice);

	HRESULT portableDeviceValues_GetStringValue(IPortableDeviceValues *pPortableDeviceValues, const PROPERTYKEY *key, PWSTR *value, DWORD *cValue);
	HRESULT portableDeviceValues_GetUnsignedIntegerValue(IPortableDeviceValues *pPortableDeviceValues, const PROPERTYKEY *key, ULONG *value);
	HRESULT portableDeviceValues_SetStringValue(IPortableDeviceValues *pPortableDeviceValues, const PROPERTYKEY *key, LPCWSTR value);
	HRESULT portableDeviceValues_SetUnsignedIntegerValue(IPortableDeviceValues *pPortableDeviceValues, const PROPERTYKEY *key, const ULONG value);

	HRESULT portableDeviceManager_GetDevices(IPortableDeviceManager *pPortableDeviceManager, PnPDeviceID **pPnPDeviceIDs, DWORD *cPnPDeviceIDs);
	HRESULT portableDeviceManager_GetDeviceFriendlyName(IPortableDeviceManager *pPortableDeviceManager, PnPDeviceID pnpDeviceID, PWSTR *pFriendlyName, DWORD *cFriendlyName);
	HRESULT portableDeviceManager_GetDeviceManufacturer(IPortableDeviceManager *pPortableDeviceManager, PnPDeviceID pnpDeviceID, PWSTR *pManufacturer, DWORD *cManufacturer);
	HRESULT portableDeviceManager_GetDeviceDescription(IPortableDeviceManager *pPortableDeviceManager, PnPDeviceID pnpDeviceID, PWSTR *pDescription, DWORD *cDescription);

	HRESULT portableDeviceContent_EnumObjects(IPortableDeviceContent *pPortableDeviceContent, DWORD flags, PWSTR parentObjectID, IPortableDeviceValues *pFilter, IEnumPortableDeviceObjectIDs **ppEnum);

	void test();

#ifdef __cplusplus
}
#endif
