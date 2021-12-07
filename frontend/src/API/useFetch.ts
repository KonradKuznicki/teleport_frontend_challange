import { useEffect, useState } from 'react';

export function useFetch<T>(url: string): {
    data: T | undefined;
    isLoading: boolean;
    isError: false | string;
} {
    const [data, setData] = useState(undefined);
    const [isLoading, setIsLoading] = useState(true);
    const [isError, setIsError] = useState<false | string>(false);

    useEffect(() => {
        const fetchData = async () => {
            setIsError(false);
            setIsLoading(true);

            try {
                const result = await fetch(url);
                if (result.status === 403) {
                    window.location.pathname = '/login';
                }

                setData(await result.json());
            } catch (error) {
                setIsError((error as Object).toString());
            }

            setIsLoading(false);
        };

        fetchData();
    }, [url]);

    return { data, isLoading, isError };
}
