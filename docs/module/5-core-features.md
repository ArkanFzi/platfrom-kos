# 5. Core Features

Detailed implementation details of the primary system functionalities.

## 1. Dynamic Room Management
Rooms are categorized and searchable. The system handles real-time status updates (Tersedia, Terisi, Booked, Perbaikan).

### Key Code: Room Filtering
```typescript
// From useHome.ts
const displayRooms = useMemo(() => {
  return realRooms.filter((room) => {
    if (selectedType !== 'all' && !room.type.includes(selectedType)) return false;
    if (room.status === 'Tidak Tersedia') return false;
    return true;
  });
}, [realRooms, selectedType]);
```

## 2. Advanced Booking Flow
The booking process is **Atomic**. It ensures that if a payment is created, the room status is updated simultaneously using database transactions.

### Atomic Transaction Example
```go
err = s.db.Transaction(func(tx *gorm.DB) error {
    // 1. Create Booking
    // 2. Setup Payment
    // 3. Update Room Status
    return nil // Commit if no error
})
```

## 3. Media Gallery
A professional, asymmetric grid showcasing the properties, optimized by Cloudinary.

### Premium UI Component
* **Asymmetric Layout**: Uses dynamic margin offsets based on data index.
* **Animated Presence**: Smooth enter/exit transitions for filtered items.
* **Vertical Bounce Loader**: An elegant "Load More" indicator.

---

> [!IMPORTANT]
> All room images must be uploaded through the Admin Dashboard to ensure they are properly processed and stored in the `koskosan/rooms` cloud folder.
