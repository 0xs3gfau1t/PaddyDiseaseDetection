import { HeatmapPoint } from '@/api/map';

function partition(
    input: { distance: number; point: HeatmapPoint }[],
    left: number,
    right: number
) {
    let pivot = input[left].distance,
        l = left + 1,
        r = right;
    while (l <= r) {
        if (input[l].distance > pivot && input[r].distance < pivot) {
            [input[l], input[r]] = [input[r], input[l]];
            l++;
            r--;
        }
        if (input[l].distance <= pivot) l++;
        if (input[r].distance >= pivot) r--;
    }
    [input[left], input[r]] = [input[r], input[left]];
    return r;
}

export function partialSort(
    input: { distance: number; point: HeatmapPoint }[],
    left: number,
    right: number,
    k: number
) {
    if (left == right) return input[left];
    let pivotIndex = partition(input, left, right);
    if (k == pivotIndex) return input[k];
    else if (k < pivotIndex) return partialSort(input, left, pivotIndex - 1, k);
    else return partialSort(input, pivotIndex + 1, right, k);
}

// https://stackoverflow.com/a/27943
export function getDistanceFromLatLonInKm(
    src: { latitude: number; longitude: number },
    dest: { latitude: number; longitude: number }
) {
    const R = 6371;
    const dLat = deg2rad(dest.latitude - src.latitude);
    const dLon = deg2rad(dest.longitude - src.longitude);
    const a =
        Math.sin(dLat / 2) * Math.sin(dLat / 2) +
        Math.cos(deg2rad(src.latitude)) *
        Math.cos(deg2rad(dest.latitude)) *
        Math.sin(dLon / 2) *
        Math.sin(dLon / 2);
    const c = 2 * Math.atan2(Math.sqrt(a), Math.sqrt(1 - a));
    return R * c * 1000;
}

function deg2rad(deg: number) {
    return deg * (Math.PI / 180);
}
