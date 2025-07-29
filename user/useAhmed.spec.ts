import { renderHook } from '@testing-library/react-hooks';
import useAhmed from './useAhmed';

describe('useAhmed', () => {
  it('should initialize correctly', () => {
    const { result } = renderHook(() => useAhmed());
    expect(result.current).toBeDefined();
  });
