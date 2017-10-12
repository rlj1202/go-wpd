#include "stdafx.h"

HRESULT portableDevice_Advise(IPortableDevice *pPortableDevice, DWORD flags, IPortableDeviceEventCallback *pCallback, IPortableDeviceValues *pParameters, PWSTR *ppCookie) {
	return pPortableDevice->Advise(flags, pCallback, pParameters, ppCookie);
}

HRESULT portableDevice_Cancel(IPortableDevice *pPortableDevice) {
	return pPortableDevice->Cancel();
}

HRESULT portableDevice_Capabilities(IPortableDevice *pPortableDevice, IPortableDeviceCapabilities **ppCapabilities) {
	return pPortableDevice->Capabilities(ppCapabilities);
}

HRESULT portableDevice_Close(IPortableDevice *pPortableDevice) {
	return pPortableDevice->Close();
}

HRESULT portableDevice_Content(IPortableDevice *pPortableDevice, IPortableDeviceContent **ppPortableDeviceContent) {
	return pPortableDevice->Content(ppPortableDeviceContent);
}

HRESULT portableDevice_GetPnPDeviceID(IPortableDevice *pPortableDevice, PnPDeviceID *pPnPDeviceID) {
	return pPortableDevice->GetPnPDeviceID(pPnPDeviceID);
}

HRESULT portableDevice_Open(IPortableDevice *pPortableDevice, PnPDeviceID pnpDeviceID, IPortableDeviceValues *pClientInfo) {
	return pPortableDevice->Open(pnpDeviceID, pClientInfo);
}

HRESULT portableDevice_SendCommand(IPortableDevice *pPortableDevice, DWORD flags, IPortableDeviceValues *pValues, IPortableDeviceValues **ppResults) {
	return pPortableDevice->SendCommand(flags, pValues, ppResults);
}

HRESULT portableDevice_Unadvise(IPortableDevice *pPortableDevice, PWSTR pCookie) {
	return pPortableDevice->Unadvise(pCookie);
}

HRESULT portableDevice_Release(IPortableDevice *pPortableDevice) {
	return pPortableDevice->Release();
}

HRESULT portableDeviceValues_GetBoolValue(IPortableDeviceValues *pPortableDeviceValues, const PROPERTYKEY *key, BOOL *value) {
	return pPortableDeviceValues->GetBoolValue(*key, value);
}

HRESULT portableDeviceValues_GetStringValue(IPortableDeviceValues *pPortableDeviceValues, const PROPERTYKEY *key, PWSTR *value, DWORD *cValue) {
	HRESULT hr = pPortableDeviceValues->GetStringValue(*key, value);
	if (SUCCEEDED(hr)) {
		*cValue = wcslen(*value);
	}

	return hr;
}

HRESULT portableDeviceValues_GetUnsignedIntegerValue(IPortableDeviceValues *pPortableDeviceValues, const PROPERTYKEY *key, ULONG *value) {
	return pPortableDeviceValues->GetUnsignedIntegerValue(*key, value);
}

HRESULT portableDeviceValues_SetGuidValue(IPortableDeviceValues *pPortableDeviceValues, const PROPERTYKEY *key, GUID *pGuid) {
	return pPortableDeviceValues->SetGuidValue(*key, *pGuid);
}

HRESULT portableDeviceValues_SetStringValue(IPortableDeviceValues *pPortableDeviceValues, const PROPERTYKEY *key, LPCWSTR value) {
	return pPortableDeviceValues->SetStringValue(*key, value);
}

HRESULT portableDeviceValues_SetUnsignedIntegerValue(IPortableDeviceValues *pPortableDeviceValues, const PROPERTYKEY *key, const ULONG value) {
	return pPortableDeviceValues->SetUnsignedIntegerValue(*key, value);
}

HRESULT portableDeviceValues_SetUnsignedLargeIntegerValue(IPortableDeviceValues *pPortableDeviceValues, const PROPERTYKEY *key, const ULONGLONG value) {
	return pPortableDeviceValues->SetUnsignedLargeIntegerValue(*key, value);
}

HRESULT portableDeviceValues_Release(IPortableDeviceValues *pPortableDeviceValues) {
	return pPortableDeviceValues->Release();
}

HRESULT portableDeviceManager_GetDevices(IPortableDeviceManager *pPortableDeviceManager, PnPDeviceID **pPnPDeviceIDs, DWORD *cPnPDeviceIDs) {
	if (pPortableDeviceManager == NULL) {
		return E_POINTER;
	}
	
	HRESULT hr = pPortableDeviceManager->GetDevices(NULL, cPnPDeviceIDs);
	if (SUCCEEDED(hr)) {
		*pPnPDeviceIDs = (PnPDeviceID*)malloc(sizeof(PnPDeviceID) * (*cPnPDeviceIDs));
		if (*pPnPDeviceIDs == NULL) {
			return E_OUTOFMEMORY;
		}
		hr = pPortableDeviceManager->GetDevices(*pPnPDeviceIDs, cPnPDeviceIDs);
	}
	
	return hr;
}

HRESULT portableDeviceManager_GetDeviceFriendlyName(IPortableDeviceManager *pPortableDeviceManager, PnPDeviceID pnpDeviceID, PWSTR *pFriendlyName, DWORD *cFriendlyName) {
	if (pPortableDeviceManager == NULL) {
		return E_POINTER;
	}

	HRESULT hr = pPortableDeviceManager->GetDeviceFriendlyName(pnpDeviceID, NULL, cFriendlyName);
	if (SUCCEEDED(hr)) {
		*pFriendlyName = (PWSTR)malloc(sizeof(WCHAR) * (*cFriendlyName));
		if (*pFriendlyName == NULL) {
			return E_OUTOFMEMORY;
		}
		hr = pPortableDeviceManager->GetDeviceFriendlyName(pnpDeviceID, *pFriendlyName, cFriendlyName);
	}

	return hr;
}

HRESULT portableDeviceManager_GetDeviceManufacturer(IPortableDeviceManager *pPortableDeviceManager, PnPDeviceID pnpDeviceID, PWSTR *pManufacturer, DWORD *pcManufacturer) {
	HRESULT hr = pPortableDeviceManager->GetDeviceManufacturer(pnpDeviceID, NULL, pcManufacturer);
	if (SUCCEEDED(hr) && *pcManufacturer > 0) {
		*pManufacturer = (PWSTR)malloc(sizeof(WCHAR) * (*pcManufacturer));
		if (*pManufacturer == NULL) {
			return E_OUTOFMEMORY;
		}
		hr = pPortableDeviceManager->GetDeviceManufacturer(pnpDeviceID, *pManufacturer, pcManufacturer);
	}

	return hr;
}

HRESULT portableDeviceManager_GetDeviceDescription(IPortableDeviceManager *pPortableDeviceManager, PnPDeviceID pnpDeviceID, PWSTR *pDescription, DWORD *cDescription) {
	HRESULT hr = pPortableDeviceManager->GetDeviceDescription(pnpDeviceID, NULL, cDescription);
	if (SUCCEEDED(hr)) {
		*pDescription = (PWSTR)malloc(sizeof(WCHAR) * (*cDescription));
		if (*pDescription == NULL) {
			return E_OUTOFMEMORY;
		}
		hr = pPortableDeviceManager->GetDeviceDescription(pnpDeviceID, *pDescription, cDescription);
	}

	return hr;
}

HRESULT portableDeviceContent_CreateObjectWithPropertiesAndData(IPortableDeviceContent *pPortableDeviceContent, IPortableDeviceValues *pValues, IStream **ppData, DWORD *pOptimalWriteBufferSize, PWSTR *ppCookie) {
	return pPortableDeviceContent->CreateObjectWithPropertiesAndData(pValues, ppData, pOptimalWriteBufferSize, ppCookie);
}

HRESULT portableDeviceContent_EnumObjects(IPortableDeviceContent *pPortableDeviceContent, DWORD flags, PWSTR parentObjectID, IPortableDeviceValues *pFilter, IEnumPortableDeviceObjectIDs **ppEnum) {
	return pPortableDeviceContent->EnumObjects(flags, parentObjectID, pFilter, ppEnum);
}

HRESULT portableDeviceContent_Properties(IPortableDeviceContent *pPortableDeviceContent, IPortableDeviceProperties **ppPortableDeviceProperties) {
	return pPortableDeviceContent->Properties(ppPortableDeviceProperties);
}

HRESULT portableDeviceContent_Transfer(IPortableDeviceContent *pPortableDeviceContent, IPortableDeviceResources **ppPortableDeviceResources) {
	return pPortableDeviceContent->Transfer(ppPortableDeviceResources);
}

HRESULT portableDeviceContent_Delete(IPortableDeviceContent *pPortableDeviceContent, DWORD options, IPortableDevicePropVariantCollection *pObjectIDs, IPortableDevicePropVariantCollection **ppResults) {
	return pPortableDeviceContent->Delete(options, pObjectIDs, ppResults);
}

HRESULT portableDeviceKeyCollection_Add(IPortableDeviceKeyCollection *pPortableDeviceKeyCollection, const PROPERTYKEY *key) {
	return pPortableDeviceKeyCollection->Add(*key);
}

HRESULT portableDeviceProperties_GetValues(IPortableDeviceProperties *pPortableDeviceProperties, PWSTR objectID, IPortableDeviceKeyCollection *pKeys, IPortableDeviceValues **ppValues) {
	return pPortableDeviceProperties->GetValues(objectID, pKeys, ppValues);
}

HRESULT portableDeviceProperties_GetPropertyAttributes(IPortableDeviceProperties *pPortableDeviceProperties, PWSTR pObjectID, const PROPERTYKEY *key, IPortableDeviceValues **ppAttributes) {
	return pPortableDeviceProperties->GetPropertyAttributes(pObjectID, *key, ppAttributes);
}

HRESULT portableDeviceProperties_SetValues(IPortableDeviceProperties *pPortableDeviceProperties, PWSTR pObjectID, IPortableDeviceValues *pValues, IPortableDeviceValues **ppResults) {
	return pPortableDeviceProperties->SetValues(pObjectID, pValues, ppResults);
}

HRESULT portableDeviceDataStream_Commit(IPortableDeviceDataStream *pPortableDeviceDataStream, DWORD dataFlags) {
	return pPortableDeviceDataStream->Commit(dataFlags);
}

HRESULT portableDeviceDataStream_GetObjectID(IPortableDeviceDataStream *pPortableDeviceDataStream, PWSTR *pObjectID) {
	return pPortableDeviceDataStream->GetObjectID(pObjectID);
}

HRESULT portableDevicePropVariantCollection_Add(IPortableDevicePropVariantCollection *pPortableDevicePropVariantCollection, const PROPVARIANT *pValue) {
	return pPortableDevicePropVariantCollection->Add(pValue);
}

HRESULT portableDevicePropVariantCollection_GetAt(IPortableDevicePropVariantCollection *pPortableDevicePropVariantCollection, DWORD index, PROPVARIANT *value) {
	return pPortableDevicePropVariantCollection->GetAt(index, value);
}

HRESULT portableDevicePropVariantCollection_GetCount(IPortableDevicePropVariantCollection *pPortableDevicePropVariantCollection, DWORD *pCounts) {
	return pPortableDevicePropVariantCollection->GetCount(pCounts);
}


HRESULT portableDeviceResources_GetStream(IPortableDeviceResources *pPortableDeviceResources, PWSTR objectID, const PROPERTYKEY *key, DWORD mode, DWORD *optimalBufferSize, IStream **ppStream) {
	return pPortableDeviceResources->GetStream(objectID, *key, mode, optimalBufferSize, ppStream);
}

HRESULT enumPortableDeviceObjectIDs_Next(IEnumPortableDeviceObjectIDs *pEnumObjectIDs, ULONG cObjects, PWSTR *pObjIDs, DWORD *pcObjIDs, ULONG *pcPetched) {
	HRESULT hr = pEnumObjectIDs->Next(cObjects, pObjIDs, pcPetched);
	if (SUCCEEDED(hr)) {
		for (int i = 0; i < *pcPetched; i++) {
			pcObjIDs[i] = wcslen(pObjIDs[i]);
		}
	}

	return hr;
}

HRESULT stream_Commit(IStream *pStream, DWORD dataFlag) {
	return pStream->Commit(dataFlag);
}

HRESULT stream_Stat(IStream *pStream, STATSTG *pStatstg, DWORD statFlags) {
	return pStream->Stat(pStatstg, statFlags);
}

HRESULT sequentialStream_Read(ISequentialStream *pSequentialStream, LPVOID pBuffer, ULONG cb, ULONG *pcbRead) {
	return pSequentialStream->Read(pBuffer, cb, pcbRead);
}

HRESULT sequentialStream_Write(ISequentialStream *pSequentialStream, LPVOID pBuffer, ULONG cb, ULONG *pcbWritten) {
	return pSequentialStream->Write(pBuffer, cb, pcbWritten);
}

HRESULT unknown_QueryInterface(IUnknown *pUnknown, const IID *piid, LPVOID *ppvObject) {
	return pUnknown->QueryInterface(*piid, ppvObject);
}
