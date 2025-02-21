package pokecache

import (
	"slices"
	"sync"
	"testing"
	"time"
)

// Check cache structure
func TestNewCache(t *testing.T) {
	testCase := struct{
        input time.Duration
        expected *Cache
    }{
        input: 5 * time.Second,
        expected: &Cache{
            data: map[string]cacheEntry{},
            mu: &sync.RWMutex{},
        },
    }

    actualCache := NewCache(testCase.input)
    if len(actualCache.data) != len(testCase.expected.data) && actualCache.mu != testCase.expected.mu {
        t.Error("Wrong cache structure")
    }
}

// check saving cache
func TestAdd(t *testing.T) {
    testData := []byte("test")
    testCase := struct{
        inputKey string
        inputVal []byte
        expected []byte
    }{
        inputKey: "https://www.test.key.com",
        inputVal: testData,
        expected: testData,
    }

    cache := NewCache(1 * time.Second)
    cache.Add(testCase.inputKey, testCase.inputVal)
    cached, exists := cache.data[testCase.inputKey]
    if !exists {
        t.Errorf("No entries available by this key: %s", testCase.inputKey)
    }

    if !slices.Equal(cached.val, testData) {
        t.Errorf("Wrong data in cache: %v\nExpected: %v", cached.val, testCase.expected)
    }
}

// check getting cache
func TestGet(t *testing.T) {
    testData := []byte("test")
    testCase := struct{
        inputKey string
        inputVal []byte
        expectedExistence bool
        expectedValue []byte
    }{
        inputKey: "https://www.test.key.com",
        inputVal: testData,
        expectedExistence: true,
        expectedValue: testData,
    }

    cache := NewCache(1 * time.Second)
    cache.Add(testCase.inputKey, testCase.inputVal)

    cached, exists := cache.Get(testCase.inputKey)
    if exists != testCase.expectedExistence {
        t.Errorf("Wrong existence state: %v\nExpected: %v", exists, testCase.expectedExistence)
    }

    if !slices.Equal(cached, testCase.expectedValue) {
        t.Errorf("Wrong data in cache: %v\nExpected: %v", cached, testCase.expectedValue)
    }
}

// check timed cache clearing
func TestReapLoop(t *testing.T) {
    testKey := "https://www.test.key.com"
    testData := []byte("test")
    testCases := []struct{
        inputKey string
        inputVal []byte
        cachedFor time.Duration
        expected bool
    }{
        {
            inputKey: testKey,
            inputVal: testData,
            cachedFor: time.Duration(5) * time.Millisecond,
            expected: false,
        },
        {
            inputKey: testKey,
            inputVal: testData,
            cachedFor: time.Duration(10) * time.Millisecond,
            expected: true,
        },
    }

    for _, test := range testCases {
        cache := NewCache(test.cachedFor)
        cache.Add(testKey, testData)
        time.Sleep(5 * time.Millisecond)

        _, exists := cache.Get(test.inputKey)
        if exists != test.expected {
            t.Errorf("Wrong existence state: %v\nExpected: %v", exists, test.expected)
        }
    }
}