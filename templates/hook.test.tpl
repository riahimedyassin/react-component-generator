import { renderHook } from '@testing-library/react-hooks';
import {{HOOK_NAME}} from './{{HOOK_NAME}}';

describe('{{HOOK_NAME}}', () => {
  it('should initialize correctly', () => {
    const { result } = renderHook(() => {{HOOK_NAME}}());
    expect(result.current).toBeDefined();
  });
